// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	model "go-common/app/service/main/spy/model"
	"go-common/library/time"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClearReliveTimesReply) DeepCopyInto(out *ClearReliveTimesReply) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClearReliveTimesReply.
func (in *ClearReliveTimesReply) DeepCopy() *ClearReliveTimesReply {
	if in == nil {
		return nil
	}
	out := new(ClearReliveTimesReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClearReliveTimesReq) DeepCopyInto(out *ClearReliveTimesReq) {
	*out = *in
	if in.Arg != nil {
		in, out := &in.Arg, &out.Arg
		*out = new(ModelArgReset)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClearReliveTimesReq.
func (in *ClearReliveTimesReq) DeepCopy() *ClearReliveTimesReq {
	if in == nil {
		return nil
	}
	out := new(ClearReliveTimesReq)
	in.DeepCopyInto(out)
	return out
}

func (in *ClearReliveTimesReq) DeepCopyAsIntoArgReset(out *model.ArgReset) {
	if out == nil {
		return
	}
	out.Mid = in.Arg.Mid
	out.BaseScore = in.Arg.BaseScore
	out.EventScore = in.Arg.EventScore
	out.Operator = in.Arg.Operator
	out.ReLiveTime = in.Arg.ReLiveTime
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandleEventReply) DeepCopyInto(out *HandleEventReply) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandleEventReply.
func (in *HandleEventReply) DeepCopy() *HandleEventReply {
	if in == nil {
		return nil
	}
	out := new(HandleEventReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HandleEventReq) DeepCopyInto(out *HandleEventReq) {
	*out = *in
	if in.EventMsg != nil {
		in, out := &in.EventMsg, &out.EventMsg
		*out = new(ModelEventMessage)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

func (in *HandleEventReq) DeepCopyAsIntoEventMessage(out *model.EventMessage) {
	if out == nil {
		return
	}
	out.Time = time.Time(in.EventMsg.Time)
	out.IP = in.EventMsg.Ip
	out.Service = in.EventMsg.Service
	out.Event = in.EventMsg.Event
	out.ActiveMid = in.EventMsg.ActiveMid
	out.TargetMid = in.EventMsg.TargetMid
	out.TargetID = in.EventMsg.TargetId
	out.Result = in.EventMsg.Result
	out.Effect = in.EventMsg.Effect
	out.RiskLevel = int8(in.EventMsg.RiskLevel)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HandleEventReq.
func (in *HandleEventReq) DeepCopy() *HandleEventReq {
	if in == nil {
		return nil
	}
	out := new(HandleEventReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelArgReset) DeepCopyInto(out *ModelArgReset) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelArgReset.
func (in *ModelArgReset) DeepCopy() *ModelArgReset {
	if in == nil {
		return nil
	}
	out := new(ModelArgReset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelEventMessage) DeepCopyInto(out *ModelEventMessage) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelEventMessage.
func (in *ModelEventMessage) DeepCopy() *ModelEventMessage {
	if in == nil {
		return nil
	}
	out := new(ModelEventMessage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelStatistics) DeepCopyInto(out *ModelStatistics) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelStatistics.
func (in *ModelStatistics) DeepCopy() *ModelStatistics {
	if in == nil {
		return nil
	}
	out := new(ModelStatistics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyAsIntoStatistics is an autogenerated deepcopy function, copying the receiver, writing into model.Statistics.
func (in *ModelStatistics) DeepCopyAsIntoStatistics(out *model.Statistics) {
	out.Quantity = in.Quantity
	out.EventName = in.EventName
	return
}

// DeepCopyFromStatistics is an autogenerated deepcopy function, copying the receiver, writing into model.Statistics.
func (out *ModelStatistics) DeepCopyFromStatistics(in *model.Statistics) {
	out.Quantity = in.Quantity
	out.EventId = in.EventID
	out.EventName = in.EventName
	return
}

// DeepCopyAsStatistics is an autogenerated deepcopy function, copying the receiver, creating a new model.Statistics.
func (in *ModelStatistics) DeepCopyAsStatistics() *model.Statistics {
	if in == nil {
		return nil
	}
	out := new(model.Statistics)
	in.DeepCopyAsIntoStatistics(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelUserInfo) DeepCopyInto(out *ModelUserInfo) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelUserInfo.
func (in *ModelUserInfo) DeepCopy() *ModelUserInfo {
	if in == nil {
		return nil
	}
	out := new(ModelUserInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyAsIntoUserInfo is an autogenerated deepcopy function, copying the receiver, writing into model.UserInfo.
func (in *ModelUserInfo) DeepCopyAsIntoUserInfo(out *model.UserInfo) {
	out.Mid = in.Mid
	return
}

// DeepCopyFromUserInfo is an autogenerated deepcopy function, copying the receiver, writing into model.UserInfo.
func (out *ModelUserInfo) DeepCopyFromUserInfo(in *model.UserInfo) {
	if in == nil {
		return
	}
	out.Mid = in.Mid
	out.Id = in.ID
	out.Score = int32(in.Score)
	out.BaseScore = int32(in.BaseScore)
	out.EventScore = int32(in.EventScore)
	out.State = int32(in.State)
	out.ReliveTimes = int32(in.ReliveTimes)
	out.Ctime = int64(in.CTime)
	out.Mtime = int64(in.MTime)
	return
}

// DeepCopyAsUserInfo is an autogenerated deepcopy function, copying the receiver, creating a new model.UserInfo.
func (in *ModelUserInfo) DeepCopyAsUserInfo() *model.UserInfo {
	if in == nil {
		return nil
	}
	out := new(model.UserInfo)
	in.DeepCopyAsIntoUserInfo(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PingReply) DeepCopyInto(out *PingReply) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PingReply.
func (in *PingReply) DeepCopy() *PingReply {
	if in == nil {
		return nil
	}
	out := new(PingReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PingReq) DeepCopyInto(out *PingReq) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PingReq.
func (in *PingReq) DeepCopy() *PingReq {
	if in == nil {
		return nil
	}
	out := new(PingReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurgeUserReply) DeepCopyInto(out *PurgeUserReply) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurgeUserReply.
func (in *PurgeUserReply) DeepCopy() *PurgeUserReply {
	if in == nil {
		return nil
	}
	out := new(PurgeUserReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurgeUserReq) DeepCopyInto(out *PurgeUserReq) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurgeUserReq.
func (in *PurgeUserReq) DeepCopy() *PurgeUserReq {
	if in == nil {
		return nil
	}
	out := new(PurgeUserReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReBuildPortraitReply) DeepCopyInto(out *ReBuildPortraitReply) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReBuildPortraitReply.
func (in *ReBuildPortraitReply) DeepCopy() *ReBuildPortraitReply {
	if in == nil {
		return nil
	}
	out := new(ReBuildPortraitReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReBuildPortraitReq) DeepCopyInto(out *ReBuildPortraitReq) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReBuildPortraitReq.
func (in *ReBuildPortraitReq) DeepCopy() *ReBuildPortraitReq {
	if in == nil {
		return nil
	}
	out := new(ReBuildPortraitReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RefreshBaseScoreReply) DeepCopyInto(out *RefreshBaseScoreReply) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RefreshBaseScoreReply.
func (in *RefreshBaseScoreReply) DeepCopy() *RefreshBaseScoreReply {
	if in == nil {
		return nil
	}
	out := new(RefreshBaseScoreReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RefreshBaseScoreReq) DeepCopyInto(out *RefreshBaseScoreReq) {
	*out = *in
	if in.Arg != nil {
		in, out := &in.Arg, &out.Arg
		*out = new(ModelArgReset)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RefreshBaseScoreReq.
func (in *RefreshBaseScoreReq) DeepCopy() *RefreshBaseScoreReq {
	if in == nil {
		return nil
	}
	out := new(RefreshBaseScoreReq)
	in.DeepCopyInto(out)
	return out
}

func (in *RefreshBaseScoreReq) DeepCopyAsIntoArgReset(out *model.ArgReset) {
	if out == nil {
		return
	}
	out.Mid = in.Arg.Mid
	out.BaseScore = in.Arg.BaseScore
	out.EventScore = in.Arg.EventScore
	out.Operator = in.Arg.Operator
	out.ReLiveTime = in.Arg.ReLiveTime
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatByIDGroupEventReply) DeepCopyInto(out *StatByIDGroupEventReply) {
	*out = *in
	if in.Res != nil {
		in, out := &in.Res, &out.Res
		*out = make([]*ModelStatistics, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ModelStatistics)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatByIDGroupEventReply.
func (in *StatByIDGroupEventReply) DeepCopy() *StatByIDGroupEventReply {
	if in == nil {
		return nil
	}
	out := new(StatByIDGroupEventReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyAsIntoStatistics is an autogenerated deepcopy function, copying the receiver, writing into model.Statistics.
func (in *StatByIDGroupEventReply) DeepCopyAsIntoStatistics(out *model.Statistics) {
	return
}

// DeepCopyFromStatistics is an autogenerated deepcopy function, copying the receiver, writing into model.Statistics.
func (out *StatByIDGroupEventReply) DeepCopyFromStatistics(in []*model.Statistics) {
	out.Res = make([]*ModelStatistics, 0, len(in))
	for _, v := range in {
		statistic := new(ModelStatistics)
		statistic.DeepCopyFromStatistics(v)
		out.Res = append(out.Res, statistic)
	}
	return
}

// DeepCopyAsStatistics is an autogenerated deepcopy function, copying the receiver, creating a new model.Statistics.
func (in *StatByIDGroupEventReply) DeepCopyAsStatistics() *model.Statistics {
	if in == nil {
		return nil
	}
	out := new(model.Statistics)
	in.DeepCopyAsIntoStatistics(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatByIDGroupEventReq) DeepCopyInto(out *StatByIDGroupEventReq) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatByIDGroupEventReq.
func (in *StatByIDGroupEventReq) DeepCopy() *StatByIDGroupEventReq {
	if in == nil {
		return nil
	}
	out := new(StatByIDGroupEventReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatByIDReply) DeepCopyInto(out *StatByIDReply) {
	*out = *in
	if in.Stat != nil {
		in, out := &in.Stat, &out.Stat
		*out = make([]*ModelStatistics, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ModelStatistics)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatByIDReply.
func (in *StatByIDReply) DeepCopy() *StatByIDReply {
	if in == nil {
		return nil
	}
	out := new(StatByIDReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyAsIntoStatistics is an autogenerated deepcopy function, copying the receiver, writing into model.Statistics.
func (in *StatByIDReply) DeepCopyAsIntoStatistics(out *model.Statistics) {
	return
}

// DeepCopyFromStatistics is an autogenerated deepcopy function, copying the receiver, writing into model.Statistics.
func (out *StatByIDReply) DeepCopyFromStatistics(in []*model.Statistics) {
	out.Stat = make([]*ModelStatistics, 0, len(in))
	for _, v := range in {
		statistic := new(ModelStatistics)
		statistic.DeepCopyFromStatistics(v)
		out.Stat = append(out.Stat, statistic)
	}
	return
}

// DeepCopyAsStatistics is an autogenerated deepcopy function, copying the receiver, creating a new model.Statistics.
func (in *StatByIDReply) DeepCopyAsStatistics() *model.Statistics {
	if in == nil {
		return nil
	}
	out := new(model.Statistics)
	in.DeepCopyAsIntoStatistics(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatByIDReq) DeepCopyInto(out *StatByIDReq) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatByIDReq.
func (in *StatByIDReq) DeepCopy() *StatByIDReq {
	if in == nil {
		return nil
	}
	out := new(StatByIDReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateBaseScoreReply) DeepCopyInto(out *UpdateBaseScoreReply) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateBaseScoreReply.
func (in *UpdateBaseScoreReply) DeepCopy() *UpdateBaseScoreReply {
	if in == nil {
		return nil
	}
	out := new(UpdateBaseScoreReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateBaseScoreReq) DeepCopyInto(out *UpdateBaseScoreReq) {
	*out = *in
	if in.Arg != nil {
		in, out := &in.Arg, &out.Arg
		*out = new(ModelArgReset)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateBaseScoreReq.
func (in *UpdateBaseScoreReq) DeepCopy() *UpdateBaseScoreReq {
	if in == nil {
		return nil
	}
	out := new(UpdateBaseScoreReq)
	in.DeepCopyInto(out)
	return out
}

func (in *UpdateBaseScoreReq) DeepCopyAsIntoArgReset(out *model.ArgReset) {
	if out == nil {
		return
	}
	out.Mid = in.Arg.Mid
	out.BaseScore = in.Arg.BaseScore
	out.EventScore = in.Arg.EventScore
	out.Operator = in.Arg.Operator
	out.ReLiveTime = in.Arg.ReLiveTime
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateEventScoreReply) DeepCopyInto(out *UpdateEventScoreReply) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateEventScoreReply.
func (in *UpdateEventScoreReply) DeepCopy() *UpdateEventScoreReply {
	if in == nil {
		return nil
	}
	out := new(UpdateEventScoreReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateEventScoreReq) DeepCopyInto(out *UpdateEventScoreReq) {
	*out = *in
	if in.Arg != nil {
		in, out := &in.Arg, &out.Arg
		*out = new(ModelArgReset)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateEventScoreReq.
func (in *UpdateEventScoreReq) DeepCopy() *UpdateEventScoreReq {
	if in == nil {
		return nil
	}
	out := new(UpdateEventScoreReq)
	in.DeepCopyInto(out)
	return out
}

func (in *UpdateEventScoreReq) DeepCopyAsIntoArgReset(out *model.ArgReset) {
	if out == nil {
		return
	}
	out.Mid = in.Arg.Mid
	out.BaseScore = in.Arg.BaseScore
	out.EventScore = in.Arg.EventScore
	out.Operator = in.Arg.Operator
	out.ReLiveTime = in.Arg.ReLiveTime
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateUserScoreReply) DeepCopyInto(out *UpdateUserScoreReply) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateUserScoreReply.
func (in *UpdateUserScoreReply) DeepCopy() *UpdateUserScoreReply {
	if in == nil {
		return nil
	}
	out := new(UpdateUserScoreReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateUserScoreReq) DeepCopyInto(out *UpdateUserScoreReq) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateUserScoreReq.
func (in *UpdateUserScoreReq) DeepCopy() *UpdateUserScoreReq {
	if in == nil {
		return nil
	}
	out := new(UpdateUserScoreReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserInfoAsynReply) DeepCopyInto(out *UserInfoAsynReply) {
	*out = *in
	if in.Ui != nil {
		in, out := &in.Ui, &out.Ui
		*out = new(ModelUserInfo)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserInfoAsynReply.
func (in *UserInfoAsynReply) DeepCopy() *UserInfoAsynReply {
	if in == nil {
		return nil
	}
	out := new(UserInfoAsynReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyAsIntoUserInfo is an autogenerated deepcopy function, copying the receiver, writing into model.UserInfo.
func (in *UserInfoAsynReply) DeepCopyAsIntoUserInfo(out *model.UserInfo) {
	return
}

// DeepCopyFromUserInfo is an autogenerated deepcopy function, copying the receiver, writing into model.UserInfo.
func (out *UserInfoAsynReply) DeepCopyFromUserInfo(in *model.UserInfo) {
	if in == nil {
		return
	}
	out.Ui = new(ModelUserInfo)
	out.Ui.DeepCopyFromUserInfo(in)
	return
}

// DeepCopyAsUserInfo is an autogenerated deepcopy function, copying the receiver, creating a new model.UserInfo.
func (in *UserInfoAsynReply) DeepCopyAsUserInfo() *model.UserInfo {
	if in == nil {
		return nil
	}
	out := new(model.UserInfo)
	in.DeepCopyAsIntoUserInfo(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserInfoAsynReq) DeepCopyInto(out *UserInfoAsynReq) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserInfoAsynReq.
func (in *UserInfoAsynReq) DeepCopy() *UserInfoAsynReq {
	if in == nil {
		return nil
	}
	out := new(UserInfoAsynReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserInfoReply) DeepCopyInto(out *UserInfoReply) {
	*out = *in
	if in.Ui != nil {
		in, out := &in.Ui, &out.Ui
		*out = new(ModelUserInfo)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserInfoReply.
func (in *UserInfoReply) DeepCopy() *UserInfoReply {
	if in == nil {
		return nil
	}
	out := new(UserInfoReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyAsIntoUserInfo is an autogenerated deepcopy function, copying the receiver, writing into model.UserInfo.
func (in *UserInfoReply) DeepCopyAsIntoUserInfo(out *model.UserInfo) {
	return
}

// DeepCopyFromUserInfo is an autogenerated deepcopy function, copying the receiver, writing into model.UserInfo.
func (out *UserInfoReply) DeepCopyFromUserInfo(in *model.UserInfo) {
	if in == nil {
		return
	}
	out.Ui = new(ModelUserInfo)
	out.Ui.DeepCopyFromUserInfo(in)
	return
}

// DeepCopyAsUserInfo is an autogenerated deepcopy function, copying the receiver, creating a new model.UserInfo.
func (in *UserInfoReply) DeepCopyAsUserInfo() *model.UserInfo {
	if in == nil {
		return nil
	}
	out := new(model.UserInfo)
	in.DeepCopyAsIntoUserInfo(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserInfoReq) DeepCopyInto(out *UserInfoReq) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserInfoReq.
func (in *UserInfoReq) DeepCopy() *UserInfoReq {
	if in == nil {
		return nil
	}
	out := new(UserInfoReq)
	in.DeepCopyInto(out)
	return out
}

func (out *InfoReply) DeepCopyFromUserInfo(in *model.UserInfo) {
	if in == nil {
		return
	}
	out.Ui = new(ModelUserInfo)
	out.Ui.DeepCopyFromUserInfo(in)
	return
}