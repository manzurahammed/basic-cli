package controllers

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/manzurahammed/rm-cli/server/models"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func ctxParam(ctx context.Context, key string) urlParam  {
	ps, ok := ctx.Value(ctxKey(paramsKey)).(map[string]urlParam)
	if !ok {
		return urlParam{}
	}
	return ps[key]
}

func parseIDParam(ctx context.Context) (int, error){
	id, err := strconv.Atoi(ctxParam(ctx,idParamName ).value)
	if err!=nil {
		return 0, models.DataValidationError{Message: "invalid id provided"}
	}
	return id,nil
}

func parseIDsParam(ctx context.Context) ([]int, error){
	idsParam := strings.Split(ctxParam(ctx,idsParamName).value,",")
	var res []int
	var Invalid []int
	for _,id := range idsParam{
		n,err := strconv.Atoi(id)
		if err!=nil {
			Invalid = append(Invalid,n)
		}
		res = append(res,n)
	}
	if len(Invalid)>0 {
		return []int{},models.DataValidationError{Message: fmt.Sprintf("invalid ids, %v",Invalid)}
	}
	return  res,nil
}