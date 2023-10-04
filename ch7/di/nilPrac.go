package main

type myErr struct {
	status int
	msg    string
}

func (me myErr) Error() string {
	return me.msg
}

// 사용자 정의 인터페이스는 타입이 이미 있기때문에 결국 제로값도 nil이 아님
//func getZeroMyError(flag bool) error {
//	var me myErr
//	if flag {
//		return me
//	}
//	me = myErr{
//		status: 111,
//		msg:    "this is error 다",
//	}
//	return me
//}

func getZeroMyError(flag bool) error {
	if flag {
		return nil
	}
	return myErr{
		status: 111,
		msg:    "this is error 다",
	}
}
