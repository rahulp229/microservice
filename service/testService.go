package service

import "microservice/executor"

type TestService interface {
	FetchData(fsyms, tsyms string) (*executor.Response, error)
}

type testService struct {
	exe executor.TestExecutor
}

func NewTestService(testExe executor.TestExecutor) TestService {
	return testService{exe: testExe}
}

func (ts testService) FetchData(fsyms, tsyms string) (*executor.Response, error) {
	resp, err := ts.exe.FetchData(fsyms, tsyms)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
