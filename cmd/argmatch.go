package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)


type ArgConverter func(string)(any, error)

type ArgMatchBuilder struct{
	argConverters []ArgConverter
}

func NewArgMatchBuilder() (*ArgMatchBuilder){
	var argConverters []ArgConverter
	return &ArgMatchBuilder{
		argConverters: argConverters,
	}
}


func (amb *ArgMatchBuilder) GetArgConverter(idx int) (ArgConverter, error){
	if numConverters := len(amb.argConverters); numConverters -1 < idx && numConverters != 1{
		return nil, errors.New(fmt.Sprintf("no converters for argument %d", idx))
	}else{
		return amb.argConverters[idx], nil
	}

}

func IntegerArgMatcher(args []string) (*ArgMatchBuilder){
	amb := NewArgMatchBuilder()
	// Utilise the fact that we use the same converter over and over again if only one is specified
	amb.WithConverter(IntegerConverter)
	return amb
}

func (amb *ArgMatchBuilder) WithConverter(converterFunc func(string)(any,error)){
	amb.argConverters = append(amb.argConverters, converterFunc)
}

func (amb *ArgMatchBuilder) Match(args []string) ([]interface{}, error) {
	convertedArgs := make([]interface{}, len(args))
	for i, arg:= range args{
		converter, err := amb.GetArgConverter(i)
		if err != nil{
			log.Fatal(err.Error())
		}
		v, convertErr := converter(arg)
		convertedArgs[i] = v
		if convertErr != nil{
			return nil, convertErr
		}
	}
	return convertedArgs, nil
}


func IntegerConverter(arg string) (any, error){
	return strconv.Atoi(arg)
}

