package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
	"gorm.io/gorm"

	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/model"
	"github.com/airren/echo-bio-backend/model/req"
	"github.com/airren/echo-bio-backend/model/vo"
	"github.com/airren/echo-bio-backend/utils"
)

func QueryAlgorithm(c context.Context, req req.AlgorithmReq) (algoVOs []*vo.AlgorithmVO, pageInfo *model.PageInfo, err error) {
	algo := AlgorithmToEntity(req)
	algos, err := dal.ListAlgorithms(c, algo, &req.PageInfo)
	for _, a := range algos {
		params, _ := dal.QueryParameter(c, &model.AlgoParameter{
			AlgorithmId: a.Id,
		})
		a.Parameters = params
		algoVOs = append(algoVOs, AlgorithmToVO(*a))
	}
	return algoVOs, &req.PageInfo, err
}

func CreateAlgorithm(c context.Context, algoReq req.AlgorithmReq) error {
	algo := AlgorithmToEntity(algoReq)
	OldAlgo, err := dal.QueryAlgorithmsByName(c, algo.Name)
	if err == nil {
		return fmt.Errorf("algo <%v> already exist", OldAlgo.Name)
	} else if err != nil && err != gorm.ErrRecordNotFound {
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

	id, _ := strconv.ParseInt(algo.Image, 10, 64)
	file, err := dal.QueryFileById(c, uint64(id))
	if err != nil {
		return err
	}
	algo.Image = file.URLPath
	_, err = dal.CreateAlgorithm(c, algo)
	for _, param := range algo.Parameters {
		param.Id = utils.GenerateId()
		param.AlgorithmId = algo.Id
		if _, err := dal.CreateParameter(c, param); err != nil {
			return err
		}
	}
	return err
}

func CreateAlgorithmByFile(c context.Context, file multipart.File) error {
	algo := &model.Algorithm{}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	if err := yaml.Unmarshal(buf.Bytes(), algo); err != nil {
		return err
	}
	OldAlgo, err := dal.QueryAlgorithmsByName(c, algo.Name)
	if err == nil {
		return fmt.Errorf("algo <%v> already exist", OldAlgo.Name)
	} else if err != nil && err != gorm.ErrRecordNotFound {
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
	if err == nil {
		return fmt.Errorf("algo <%v> already exist", OldAlgo.Name)
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("query old record failed %v", err)
	}
	algo.Id = OldAlgo.Id
	_, err = dal.UpdateAlgorithm(c, algo)
	return err
}

func AlgorithmToEntity(req req.AlgorithmReq) *model.Algorithm {
	var Id int64
	if req.Id != "" {
		Id, _ = strconv.ParseInt(req.Id, 10, 64)
	}
	return &model.Algorithm{
		RecordMeta:  model.RecordMeta{Id: uint64(Id)},
		Name:        req.Name,
		Label:       req.Label,
		Image:       req.Image,
		Description: req.Description,
		Point:       req.Point,
		Favourite:   req.Favourite,
		Parameters:  req.Parameters,
		Command:     req.Command,
		DockerImage: req.DockerImage,
		Document:    req.Document,
		GroupId:     req.Group,
	}
}

func AlgorithmToVO(algorithm model.Algorithm) *vo.AlgorithmVO {
	return &vo.AlgorithmVO{
		RecordMeta:  RecordMetaToVO(algorithm.RecordMeta),
		Name:        algorithm.Name,
		Label:       algorithm.Label,
		Image:       algorithm.Image,
		Description: algorithm.Description,
		Point:       algorithm.Point,
		Favourite:   algorithm.Favourite,
		Parameters:  algorithm.Parameters,
		Command:     algorithm.Command,
		DockerImage: algorithm.DockerImage,
		Document:    algorithm.Document,
		GroupId:     algorithm.GroupId,
	}
}
