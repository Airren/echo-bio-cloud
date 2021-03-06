package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/model/vo"
	"github.com/airren/echo-bio-backend/utils"
)

func QueryAlgorithm(c context.Context, req req.AlgorithmReq) (algoVOs []*vo.AlgorithmVO, err error) {
	algo := AlgorithmToEntity(req)
	algos, err := dal.QueryAlgorithms(c, algo, &req.PageInfo)
	for _, a := range algos {
		params, _ := dal.QueryParameter(c, &model.Parameter{
			AlgorithmId: a.Id,
		})
		a.Parameters = params
		algoVOs = append(algoVOs, AlgorithmToVO(*a))
	}
	return algoVOs, err
}

func CreateAlgorithm(c context.Context, file multipart.File) error {
	algo := &model.Algorithm{}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	if err := yaml.Unmarshal(buf.Bytes(), algo); err != nil {
		return err
	}
	OldAlgo, err := dal.QueryAlgorithmsByName(c, algo.Name)
	if len(OldAlgo) > 0 {
		return fmt.Errorf("algo <%v> already exist", OldAlgo[0].Name)
	} else if err != nil {
		return fmt.Errorf("query old record failed %v", err)
	}
	userId, _ := utils.GetUserId(c)
	algo.RecordMeta = model.RecordMeta{
		Id:        utils.GenerateId(),
		AccountId: userId,
		Org:       "",
		CreatedAt: time.Now(),
		CreatedBy: "Echo Bio",
	}
	_, err = dal.CreateAlgorithm(c, algo)
	for _, param := range algo.Parameters {
		param.AlgorithmId = algo.Id
		if _, err := dal.CreateParameter(c, param); err != nil {
			return err
		}
	}
	return err
}

func UpdateAlgorithm(c context.Context, file multipart.File) error {
	algo := &model.Algorithm{}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	if err := yaml.Unmarshal(buf.Bytes(), algo); err != nil {
		return err
	}
	OldAlgo, err := dal.QueryAlgorithmsByName(c, algo.Name)
	if len(OldAlgo) == 0 {
		return fmt.Errorf("algo <%v> does not exist, please create first", OldAlgo[0].Name)
	} else if err != nil {
		return fmt.Errorf("query old record failed %v", err)
	}
	algo.Id = OldAlgo[0].Id
	_, err = dal.UpdateAlgorithm(c, algo)
	return err
}

func AlgorithmToEntity(req req.AlgorithmReq) *model.Algorithm {
	return &model.Algorithm{
		RecordMeta:  model.RecordMeta{},
		Name:        "",
		Label:       req.Label,
		Image:       "",
		Description: "",
		Price:       0,
		Favourite:   0,
		//Category:    model.Category{},
	}
}

func AlgorithmToVO(algorithm model.Algorithm) *vo.AlgorithmVO {
	return &vo.AlgorithmVO{
		RecordMeta:  algorithm.RecordMeta,
		Name:        algorithm.Name,
		Label:       algorithm.Label,
		Image:       algorithm.Image,
		Description: algorithm.Description,
		Price:       algorithm.Price,
		Favourite:   algorithm.Favourite,
		Parameters:  algorithm.Parameters,
	}
}
