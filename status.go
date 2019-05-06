package libra

import (
	"fmt"
)

/*
StatusCode shows Status
*/
type StatusCode int

//StatusCodes
const (
	NG  StatusCode = iota //Unknown
	OK                    //No Problem
	WA                    //Wrong Answer
	RE                    //Runtime Error
	CE                    //Compile Error
	TLE                   //Time Limit Exceeded
	IE                    //Internal Error
)

func (code StatusCode) String() string {
	switch code {
	case NG:
		return "NG"
	case OK:
		return "OK"
	case WA:
		return "Wrong Answer"
	case RE:
		return "Runtime Error"
	case CE:
		return "Compile Error"
	case TLE:
		return "Time Limit Exceeded"
	case IE:
		return "Internal Error"
	default:
		return "Unknown"
	}
}

/*
Status shows result of program
*/
type Status struct {
	Code StatusCode
	Msg  string
}

func (status Status) String() string {
	return fmt.Sprintf("[Code:%v, Msg: %v]", status.Code, status.Msg)
}
