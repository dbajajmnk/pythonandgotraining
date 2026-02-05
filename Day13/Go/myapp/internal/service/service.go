package service

import (
    "fmt"
    "myapp/pkg/util"
)

type Service struct {
    appName string
}

func NewService(appName string) *Service {
    return &Service{appName: appName}
}

func (s *Service) AddAndDescribe(a int, b int) string {
    sum := util.Add(a, b)
    isEven := util.IsEven(sum)

    return fmt.Sprintf("[%s] sum(%d,%d)=%d | even=%t", s.appName, a, b, sum, isEven)
}
