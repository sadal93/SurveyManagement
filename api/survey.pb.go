// Code generated by protoc-gen-go. DO NOT EDIT.
// source: survey.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Question struct {
	Id                   string                           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string                           `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	QuestionWithLanguage []*Question_QuestionWithLanguage `protobuf:"bytes,3,rep,name=questionWithLanguage,proto3" json:"questionWithLanguage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *Question) Reset()         { *m = Question{} }
func (m *Question) String() string { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()    {}
func (*Question) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{0}
}

func (m *Question) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question.Unmarshal(m, b)
}
func (m *Question) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question.Marshal(b, m, deterministic)
}
func (m *Question) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question.Merge(m, src)
}
func (m *Question) XXX_Size() int {
	return xxx_messageInfo_Question.Size(m)
}
func (m *Question) XXX_DiscardUnknown() {
	xxx_messageInfo_Question.DiscardUnknown(m)
}

var xxx_messageInfo_Question proto.InternalMessageInfo

func (m *Question) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Question) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Question) GetQuestionWithLanguage() []*Question_QuestionWithLanguage {
	if m != nil {
		return m.QuestionWithLanguage
	}
	return nil
}

type Question_AnswerOptions struct {
	OpenOptions          *Question_OpenOptions     `protobuf:"bytes,1,opt,name=openOptions,proto3" json:"openOptions,omitempty"`
	MultipleOptions      *Question_MultipleOptions `protobuf:"bytes,2,opt,name=multipleOptions,proto3" json:"multipleOptions,omitempty"`
	MatrixOptions        []*Question_MatrixOptions `protobuf:"bytes,3,rep,name=matrixOptions,proto3" json:"matrixOptions,omitempty"`
	HideQuestion         []*Question_HideQuestion  `protobuf:"bytes,4,rep,name=hideQuestion,proto3" json:"hideQuestion,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Question_AnswerOptions) Reset()         { *m = Question_AnswerOptions{} }
func (m *Question_AnswerOptions) String() string { return proto.CompactTextString(m) }
func (*Question_AnswerOptions) ProtoMessage()    {}
func (*Question_AnswerOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{0, 0}
}

func (m *Question_AnswerOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question_AnswerOptions.Unmarshal(m, b)
}
func (m *Question_AnswerOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question_AnswerOptions.Marshal(b, m, deterministic)
}
func (m *Question_AnswerOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question_AnswerOptions.Merge(m, src)
}
func (m *Question_AnswerOptions) XXX_Size() int {
	return xxx_messageInfo_Question_AnswerOptions.Size(m)
}
func (m *Question_AnswerOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Question_AnswerOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Question_AnswerOptions proto.InternalMessageInfo

func (m *Question_AnswerOptions) GetOpenOptions() *Question_OpenOptions {
	if m != nil {
		return m.OpenOptions
	}
	return nil
}

func (m *Question_AnswerOptions) GetMultipleOptions() *Question_MultipleOptions {
	if m != nil {
		return m.MultipleOptions
	}
	return nil
}

func (m *Question_AnswerOptions) GetMatrixOptions() []*Question_MatrixOptions {
	if m != nil {
		return m.MatrixOptions
	}
	return nil
}

func (m *Question_AnswerOptions) GetHideQuestion() []*Question_HideQuestion {
	if m != nil {
		return m.HideQuestion
	}
	return nil
}

type Question_OpenOptions struct {
	DataType             string   `protobuf:"bytes,1,opt,name=dataType,proto3" json:"dataType,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Question_OpenOptions) Reset()         { *m = Question_OpenOptions{} }
func (m *Question_OpenOptions) String() string { return proto.CompactTextString(m) }
func (*Question_OpenOptions) ProtoMessage()    {}
func (*Question_OpenOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{0, 1}
}

func (m *Question_OpenOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question_OpenOptions.Unmarshal(m, b)
}
func (m *Question_OpenOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question_OpenOptions.Marshal(b, m, deterministic)
}
func (m *Question_OpenOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question_OpenOptions.Merge(m, src)
}
func (m *Question_OpenOptions) XXX_Size() int {
	return xxx_messageInfo_Question_OpenOptions.Size(m)
}
func (m *Question_OpenOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Question_OpenOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Question_OpenOptions proto.InternalMessageInfo

func (m *Question_OpenOptions) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *Question_OpenOptions) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Question_OpenOptions) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type Question_MultipleOptions struct {
	MultipleOptions      []string `protobuf:"bytes,1,rep,name=multipleOptions,proto3" json:"multipleOptions,omitempty"`
	SelectedOption       string   `protobuf:"bytes,2,opt,name=selectedOption,proto3" json:"selectedOption,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Question_MultipleOptions) Reset()         { *m = Question_MultipleOptions{} }
func (m *Question_MultipleOptions) String() string { return proto.CompactTextString(m) }
func (*Question_MultipleOptions) ProtoMessage()    {}
func (*Question_MultipleOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{0, 2}
}

func (m *Question_MultipleOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question_MultipleOptions.Unmarshal(m, b)
}
func (m *Question_MultipleOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question_MultipleOptions.Marshal(b, m, deterministic)
}
func (m *Question_MultipleOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question_MultipleOptions.Merge(m, src)
}
func (m *Question_MultipleOptions) XXX_Size() int {
	return xxx_messageInfo_Question_MultipleOptions.Size(m)
}
func (m *Question_MultipleOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Question_MultipleOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Question_MultipleOptions proto.InternalMessageInfo

func (m *Question_MultipleOptions) GetMultipleOptions() []string {
	if m != nil {
		return m.MultipleOptions
	}
	return nil
}

func (m *Question_MultipleOptions) GetSelectedOption() string {
	if m != nil {
		return m.SelectedOption
	}
	return ""
}

type Question_MatrixOptions struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Options              []string `protobuf:"bytes,2,rep,name=options,proto3" json:"options,omitempty"`
	SelectedOption       string   `protobuf:"bytes,3,opt,name=selectedOption,proto3" json:"selectedOption,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Question_MatrixOptions) Reset()         { *m = Question_MatrixOptions{} }
func (m *Question_MatrixOptions) String() string { return proto.CompactTextString(m) }
func (*Question_MatrixOptions) ProtoMessage()    {}
func (*Question_MatrixOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{0, 3}
}

func (m *Question_MatrixOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question_MatrixOptions.Unmarshal(m, b)
}
func (m *Question_MatrixOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question_MatrixOptions.Marshal(b, m, deterministic)
}
func (m *Question_MatrixOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question_MatrixOptions.Merge(m, src)
}
func (m *Question_MatrixOptions) XXX_Size() int {
	return xxx_messageInfo_Question_MatrixOptions.Size(m)
}
func (m *Question_MatrixOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Question_MatrixOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Question_MatrixOptions proto.InternalMessageInfo

func (m *Question_MatrixOptions) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Question_MatrixOptions) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Question_MatrixOptions) GetSelectedOption() string {
	if m != nil {
		return m.SelectedOption
	}
	return ""
}

type Question_HideQuestion struct {
	SelectedOption       string   `protobuf:"bytes,1,opt,name=selectedOption,proto3" json:"selectedOption,omitempty"`
	Question             string   `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Question_HideQuestion) Reset()         { *m = Question_HideQuestion{} }
func (m *Question_HideQuestion) String() string { return proto.CompactTextString(m) }
func (*Question_HideQuestion) ProtoMessage()    {}
func (*Question_HideQuestion) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{0, 4}
}

func (m *Question_HideQuestion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question_HideQuestion.Unmarshal(m, b)
}
func (m *Question_HideQuestion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question_HideQuestion.Marshal(b, m, deterministic)
}
func (m *Question_HideQuestion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question_HideQuestion.Merge(m, src)
}
func (m *Question_HideQuestion) XXX_Size() int {
	return xxx_messageInfo_Question_HideQuestion.Size(m)
}
func (m *Question_HideQuestion) XXX_DiscardUnknown() {
	xxx_messageInfo_Question_HideQuestion.DiscardUnknown(m)
}

var xxx_messageInfo_Question_HideQuestion proto.InternalMessageInfo

func (m *Question_HideQuestion) GetSelectedOption() string {
	if m != nil {
		return m.SelectedOption
	}
	return ""
}

func (m *Question_HideQuestion) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

type Question_QuestionWithLanguage struct {
	Language             string                  `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Text                 string                  `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	AnswerOptions        *Question_AnswerOptions `protobuf:"bytes,3,opt,name=answerOptions,proto3" json:"answerOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Question_QuestionWithLanguage) Reset()         { *m = Question_QuestionWithLanguage{} }
func (m *Question_QuestionWithLanguage) String() string { return proto.CompactTextString(m) }
func (*Question_QuestionWithLanguage) ProtoMessage()    {}
func (*Question_QuestionWithLanguage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{0, 5}
}

func (m *Question_QuestionWithLanguage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question_QuestionWithLanguage.Unmarshal(m, b)
}
func (m *Question_QuestionWithLanguage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question_QuestionWithLanguage.Marshal(b, m, deterministic)
}
func (m *Question_QuestionWithLanguage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question_QuestionWithLanguage.Merge(m, src)
}
func (m *Question_QuestionWithLanguage) XXX_Size() int {
	return xxx_messageInfo_Question_QuestionWithLanguage.Size(m)
}
func (m *Question_QuestionWithLanguage) XXX_DiscardUnknown() {
	xxx_messageInfo_Question_QuestionWithLanguage.DiscardUnknown(m)
}

var xxx_messageInfo_Question_QuestionWithLanguage proto.InternalMessageInfo

func (m *Question_QuestionWithLanguage) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Question_QuestionWithLanguage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Question_QuestionWithLanguage) GetAnswerOptions() *Question_AnswerOptions {
	if m != nil {
		return m.AnswerOptions
	}
	return nil
}

type QuestionID struct {
	QuestionID           string   `protobuf:"bytes,1,opt,name=questionID,proto3" json:"questionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuestionID) Reset()         { *m = QuestionID{} }
func (m *QuestionID) String() string { return proto.CompactTextString(m) }
func (*QuestionID) ProtoMessage()    {}
func (*QuestionID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{1}
}

func (m *QuestionID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestionID.Unmarshal(m, b)
}
func (m *QuestionID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestionID.Marshal(b, m, deterministic)
}
func (m *QuestionID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionID.Merge(m, src)
}
func (m *QuestionID) XXX_Size() int {
	return xxx_messageInfo_QuestionID.Size(m)
}
func (m *QuestionID) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionID.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionID proto.InternalMessageInfo

func (m *QuestionID) GetQuestionID() string {
	if m != nil {
		return m.QuestionID
	}
	return ""
}

type SurveyID struct {
	SurveyID             string   `protobuf:"bytes,1,opt,name=surveyID,proto3" json:"surveyID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SurveyID) Reset()         { *m = SurveyID{} }
func (m *SurveyID) String() string { return proto.CompactTextString(m) }
func (*SurveyID) ProtoMessage()    {}
func (*SurveyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{2}
}

func (m *SurveyID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyID.Unmarshal(m, b)
}
func (m *SurveyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyID.Marshal(b, m, deterministic)
}
func (m *SurveyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyID.Merge(m, src)
}
func (m *SurveyID) XXX_Size() int {
	return xxx_messageInfo_SurveyID.Size(m)
}
func (m *SurveyID) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyID.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyID proto.InternalMessageInfo

func (m *SurveyID) GetSurveyID() string {
	if m != nil {
		return m.SurveyID
	}
	return ""
}

type QuestionArray struct {
	Questions            []*Question `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *QuestionArray) Reset()         { *m = QuestionArray{} }
func (m *QuestionArray) String() string { return proto.CompactTextString(m) }
func (*QuestionArray) ProtoMessage()    {}
func (*QuestionArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{3}
}

func (m *QuestionArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestionArray.Unmarshal(m, b)
}
func (m *QuestionArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestionArray.Marshal(b, m, deterministic)
}
func (m *QuestionArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionArray.Merge(m, src)
}
func (m *QuestionArray) XXX_Size() int {
	return xxx_messageInfo_QuestionArray.Size(m)
}
func (m *QuestionArray) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionArray.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionArray proto.InternalMessageInfo

func (m *QuestionArray) GetQuestions() []*Question {
	if m != nil {
		return m.Questions
	}
	return nil
}

type SurveyArray struct {
	Surveys              []*SurveyData `protobuf:"bytes,1,rep,name=surveys,proto3" json:"surveys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SurveyArray) Reset()         { *m = SurveyArray{} }
func (m *SurveyArray) String() string { return proto.CompactTextString(m) }
func (*SurveyArray) ProtoMessage()    {}
func (*SurveyArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{4}
}

func (m *SurveyArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyArray.Unmarshal(m, b)
}
func (m *SurveyArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyArray.Marshal(b, m, deterministic)
}
func (m *SurveyArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyArray.Merge(m, src)
}
func (m *SurveyArray) XXX_Size() int {
	return xxx_messageInfo_SurveyArray.Size(m)
}
func (m *SurveyArray) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyArray.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyArray proto.InternalMessageInfo

func (m *SurveyArray) GetSurveys() []*SurveyData {
	if m != nil {
		return m.Surveys
	}
	return nil
}

type EmptySurvey struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptySurvey) Reset()         { *m = EmptySurvey{} }
func (m *EmptySurvey) String() string { return proto.CompactTextString(m) }
func (*EmptySurvey) ProtoMessage()    {}
func (*EmptySurvey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{5}
}

func (m *EmptySurvey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptySurvey.Unmarshal(m, b)
}
func (m *EmptySurvey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptySurvey.Marshal(b, m, deterministic)
}
func (m *EmptySurvey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptySurvey.Merge(m, src)
}
func (m *EmptySurvey) XXX_Size() int {
	return xxx_messageInfo_EmptySurvey.Size(m)
}
func (m *EmptySurvey) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptySurvey.DiscardUnknown(m)
}

var xxx_messageInfo_EmptySurvey proto.InternalMessageInfo

type SurveyData struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Questions            []*Question `protobuf:"bytes,3,rep,name=questions,proto3" json:"questions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SurveyData) Reset()         { *m = SurveyData{} }
func (m *SurveyData) String() string { return proto.CompactTextString(m) }
func (*SurveyData) ProtoMessage()    {}
func (*SurveyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{6}
}

func (m *SurveyData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyData.Unmarshal(m, b)
}
func (m *SurveyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyData.Marshal(b, m, deterministic)
}
func (m *SurveyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyData.Merge(m, src)
}
func (m *SurveyData) XXX_Size() int {
	return xxx_messageInfo_SurveyData.Size(m)
}
func (m *SurveyData) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyData.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyData proto.InternalMessageInfo

func (m *SurveyData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SurveyData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SurveyData) GetQuestions() []*Question {
	if m != nil {
		return m.Questions
	}
	return nil
}

func init() {
	proto.RegisterType((*Question)(nil), "api.Question")
	proto.RegisterType((*Question_AnswerOptions)(nil), "api.Question.AnswerOptions")
	proto.RegisterType((*Question_OpenOptions)(nil), "api.Question.OpenOptions")
	proto.RegisterType((*Question_MultipleOptions)(nil), "api.Question.MultipleOptions")
	proto.RegisterType((*Question_MatrixOptions)(nil), "api.Question.MatrixOptions")
	proto.RegisterType((*Question_HideQuestion)(nil), "api.Question.HideQuestion")
	proto.RegisterType((*Question_QuestionWithLanguage)(nil), "api.Question.QuestionWithLanguage")
	proto.RegisterType((*QuestionID)(nil), "api.QuestionID")
	proto.RegisterType((*SurveyID)(nil), "api.SurveyID")
	proto.RegisterType((*QuestionArray)(nil), "api.QuestionArray")
	proto.RegisterType((*SurveyArray)(nil), "api.SurveyArray")
	proto.RegisterType((*EmptySurvey)(nil), "api.EmptySurvey")
	proto.RegisterType((*SurveyData)(nil), "api.SurveyData")
}

func init() { proto.RegisterFile("survey.proto", fileDescriptor_a40f94eaa8e6ca46) }

var fileDescriptor_a40f94eaa8e6ca46 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0x4f, 0x8f, 0xd3, 0x3e,
	0x10, 0x55, 0x92, 0xdd, 0x6d, 0x3b, 0x69, 0xda, 0x9f, 0xac, 0x1e, 0xf2, 0x0b, 0x02, 0x55, 0x39,
	0xac, 0x8a, 0x80, 0x82, 0xda, 0x03, 0x48, 0x20, 0xa4, 0x8a, 0xa2, 0x52, 0x09, 0xb4, 0xc2, 0xfc,
	0x3b, 0x9b, 0xc6, 0xda, 0xb5, 0x48, 0xdb, 0x6c, 0xe2, 0x2c, 0xdb, 0x8f, 0xc0, 0x37, 0x40, 0x9c,
	0xf9, 0xa0, 0x28, 0x76, 0x9c, 0xda, 0x49, 0xe0, 0xe6, 0x19, 0xbf, 0xf7, 0x66, 0x9e, 0x67, 0xda,
	0x40, 0x3f, 0xcb, 0xd3, 0x1b, 0x7a, 0x98, 0x26, 0xe9, 0x9e, 0xef, 0x91, 0x43, 0x12, 0x16, 0xfe,
	0xee, 0x40, 0xf7, 0x7d, 0x4e, 0x33, 0xce, 0xf6, 0x3b, 0x34, 0x00, 0x9b, 0x45, 0xbe, 0x35, 0xb6,
	0x26, 0x3d, 0x6c, 0xb3, 0x08, 0x21, 0x38, 0xe1, 0x87, 0x84, 0xfa, 0xb6, 0xc8, 0x88, 0x33, 0xfa,
	0x0c, 0xa3, 0xeb, 0x12, 0xff, 0x85, 0xf1, 0xab, 0xb7, 0x64, 0x77, 0x99, 0x93, 0x4b, 0xea, 0x3b,
	0x63, 0x67, 0xe2, 0xce, 0xc2, 0x29, 0x49, 0xd8, 0x54, 0x09, 0x56, 0x07, 0x1d, 0x89, 0x5b, 0xf9,
	0xc1, 0x4f, 0x1b, 0xbc, 0xc5, 0x2e, 0xfb, 0x4e, 0xd3, 0x8b, 0xa4, 0xb8, 0xcc, 0xd0, 0x73, 0x70,
	0xf7, 0x09, 0xdd, 0x95, 0xa1, 0x68, 0xcb, 0x9d, 0xfd, 0x6f, 0x16, 0xb8, 0x38, 0x02, 0xb0, 0x8e,
	0x46, 0x2b, 0x18, 0x6e, 0xf3, 0x98, 0xb3, 0x24, 0xa6, 0x4a, 0xc0, 0x16, 0x02, 0x77, 0x4d, 0x81,
	0x77, 0x26, 0x08, 0xd7, 0x59, 0x68, 0x01, 0xde, 0x96, 0xf0, 0x94, 0xdd, 0x2a, 0x19, 0x69, 0xf4,
	0x4e, 0x4d, 0x46, 0x87, 0x60, 0x93, 0x81, 0x5e, 0x42, 0xff, 0x8a, 0x45, 0x54, 0x81, 0xfd, 0x13,
	0xa1, 0x10, 0x98, 0x0a, 0x6f, 0x34, 0x04, 0x36, 0xf0, 0xc1, 0x27, 0x70, 0x35, 0x9f, 0x28, 0x80,
	0x6e, 0x44, 0x38, 0xf9, 0x58, 0x4c, 0x46, 0xce, 0xaa, 0x8a, 0xd1, 0x08, 0x4e, 0x6f, 0x48, 0x9c,
	0xab, 0x91, 0xc9, 0xa0, 0xc8, 0xc6, 0x6c, 0xcb, 0xb8, 0xef, 0x8c, 0xad, 0xc9, 0x29, 0x96, 0x41,
	0xb0, 0x81, 0x61, 0xcd, 0x3d, 0x9a, 0x34, 0x5f, 0xcd, 0x1a, 0x3b, 0x93, 0x5e, 0xf3, 0x59, 0xce,
	0x61, 0x90, 0xd1, 0x98, 0x6e, 0x38, 0x8d, 0x64, 0xaa, 0xac, 0x58, 0xcb, 0x06, 0x14, 0x3c, 0xe3,
	0x6d, 0xc4, 0x4e, 0xd1, 0x5b, 0x5e, 0x76, 0x2e, 0xce, 0xc8, 0x87, 0xce, 0xbe, 0x1a, 0x52, 0x51,
	0x4e, 0x85, 0x2d, 0x65, 0x9c, 0xd6, 0x32, 0x18, 0xfa, 0xfa, 0x03, 0xb6, 0xf0, 0xac, 0x36, 0x5e,
	0xf1, 0x96, 0x6a, 0x1b, 0x4b, 0x03, 0x55, 0x1c, 0xfc, 0xb0, 0x60, 0xd4, 0xb6, 0xc0, 0x05, 0x29,
	0x56, 0x6b, 0x5f, 0x0e, 0x40, 0xc5, 0x95, 0x3d, 0x5b, 0xb3, 0xb7, 0x00, 0x8f, 0xe8, 0x9b, 0x2d,
	0x3c, 0x34, 0x56, 0xc8, 0x58, 0x7e, 0x6c, 0x32, 0xc2, 0x87, 0x00, 0x0a, 0xb8, 0x5e, 0xa2, 0x7b,
	0x00, 0xd7, 0x55, 0x54, 0xb6, 0xa0, 0x65, 0xc2, 0x73, 0xe8, 0x7e, 0x10, 0xbf, 0xf4, 0xf5, 0xb2,
	0x68, 0x36, 0x2b, 0xcf, 0xaa, 0x59, 0x15, 0x87, 0x2f, 0xc0, 0x53, 0xaa, 0x8b, 0x34, 0x25, 0x07,
	0xf4, 0x00, 0x7a, 0x4a, 0x46, 0x4e, 0xde, 0x9d, 0x79, 0x46, 0x97, 0xf8, 0x78, 0x1f, 0x3e, 0x03,
	0x57, 0x56, 0x91, 0xdc, 0xfb, 0xd0, 0x91, 0xc2, 0x8a, 0x39, 0x14, 0x4c, 0x09, 0x59, 0x12, 0x4e,
	0xb0, 0xba, 0x0f, 0x3d, 0x70, 0x5f, 0x6f, 0x13, 0x7e, 0x90, 0x77, 0xe1, 0x37, 0x80, 0x23, 0xaa,
	0xf1, 0x27, 0x34, 0x06, 0x37, 0xa2, 0xd9, 0x26, 0x65, 0xfa, 0x9a, 0xe9, 0x29, 0xb3, 0x6b, 0xe7,
	0xdf, 0x5d, 0xcf, 0x7e, 0x39, 0x70, 0x26, 0xab, 0xa1, 0x27, 0xd0, 0x7f, 0x95, 0x52, 0xc2, 0x69,
	0x19, 0xd7, 0x1b, 0x0e, 0xea, 0x09, 0xf4, 0x18, 0xfa, 0x4b, 0x1a, 0xd3, 0x8a, 0xe1, 0x69, 0x80,
	0xf5, 0x32, 0xf8, 0x4f, 0x84, 0x9a, 0xb5, 0xa2, 0xb5, 0x15, 0xe5, 0xed, 0xe8, 0x86, 0xfa, 0x1c,
	0xbc, 0x15, 0xe5, 0x8b, 0x38, 0x96, 0xb9, 0x0c, 0x35, 0xf4, 0xca, 0x0a, 0xfa, 0xb3, 0x4f, 0x61,
	0x20, 0x4d, 0x54, 0xbb, 0x6f, 0x7a, 0x0f, 0xcc, 0x10, 0xcd, 0x61, 0x20, 0x2d, 0x54, 0x99, 0xa1,
	0x01, 0x68, 0xb5, 0xf1, 0x08, 0xdc, 0x15, 0xe5, 0x7f, 0x67, 0xd4, 0x6a, 0x3c, 0x85, 0xa1, 0x34,
	0xa2, 0x32, 0x6d, 0x56, 0x90, 0xc1, 0x11, 0x66, 0xbe, 0x9e, 0x89, 0x2f, 0xd3, 0xfc, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x6a, 0x10, 0x4d, 0xd1, 0xa9, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SurveyClient is the client API for Survey service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SurveyClient interface {
	CreateSurvey(ctx context.Context, in *SurveyData, opts ...grpc.CallOption) (*SurveyData, error)
	DeleteSurvey(ctx context.Context, in *SurveyID, opts ...grpc.CallOption) (*EmptySurvey, error)
	GetSurvey(ctx context.Context, in *SurveyID, opts ...grpc.CallOption) (*SurveyData, error)
	GetAllSurveys(ctx context.Context, in *EmptySurvey, opts ...grpc.CallOption) (*SurveyArray, error)
	CreateQuestion(ctx context.Context, in *Question, opts ...grpc.CallOption) (*Question, error)
	DeleteQuestion(ctx context.Context, in *QuestionID, opts ...grpc.CallOption) (*EmptySurvey, error)
	GetQuestion(ctx context.Context, in *QuestionID, opts ...grpc.CallOption) (*Question, error)
	GetAllQuestions(ctx context.Context, in *EmptySurvey, opts ...grpc.CallOption) (*QuestionArray, error)
}

type surveyClient struct {
	cc *grpc.ClientConn
}

func NewSurveyClient(cc *grpc.ClientConn) SurveyClient {
	return &surveyClient{cc}
}

func (c *surveyClient) CreateSurvey(ctx context.Context, in *SurveyData, opts ...grpc.CallOption) (*SurveyData, error) {
	out := new(SurveyData)
	err := c.cc.Invoke(ctx, "/api.Survey/CreateSurvey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surveyClient) DeleteSurvey(ctx context.Context, in *SurveyID, opts ...grpc.CallOption) (*EmptySurvey, error) {
	out := new(EmptySurvey)
	err := c.cc.Invoke(ctx, "/api.Survey/DeleteSurvey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surveyClient) GetSurvey(ctx context.Context, in *SurveyID, opts ...grpc.CallOption) (*SurveyData, error) {
	out := new(SurveyData)
	err := c.cc.Invoke(ctx, "/api.Survey/GetSurvey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surveyClient) GetAllSurveys(ctx context.Context, in *EmptySurvey, opts ...grpc.CallOption) (*SurveyArray, error) {
	out := new(SurveyArray)
	err := c.cc.Invoke(ctx, "/api.Survey/GetAllSurveys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surveyClient) CreateQuestion(ctx context.Context, in *Question, opts ...grpc.CallOption) (*Question, error) {
	out := new(Question)
	err := c.cc.Invoke(ctx, "/api.Survey/CreateQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surveyClient) DeleteQuestion(ctx context.Context, in *QuestionID, opts ...grpc.CallOption) (*EmptySurvey, error) {
	out := new(EmptySurvey)
	err := c.cc.Invoke(ctx, "/api.Survey/DeleteQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surveyClient) GetQuestion(ctx context.Context, in *QuestionID, opts ...grpc.CallOption) (*Question, error) {
	out := new(Question)
	err := c.cc.Invoke(ctx, "/api.Survey/GetQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surveyClient) GetAllQuestions(ctx context.Context, in *EmptySurvey, opts ...grpc.CallOption) (*QuestionArray, error) {
	out := new(QuestionArray)
	err := c.cc.Invoke(ctx, "/api.Survey/GetAllQuestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SurveyServer is the server API for Survey service.
type SurveyServer interface {
	CreateSurvey(context.Context, *SurveyData) (*SurveyData, error)
	DeleteSurvey(context.Context, *SurveyID) (*EmptySurvey, error)
	GetSurvey(context.Context, *SurveyID) (*SurveyData, error)
	GetAllSurveys(context.Context, *EmptySurvey) (*SurveyArray, error)
	CreateQuestion(context.Context, *Question) (*Question, error)
	DeleteQuestion(context.Context, *QuestionID) (*EmptySurvey, error)
	GetQuestion(context.Context, *QuestionID) (*Question, error)
	GetAllQuestions(context.Context, *EmptySurvey) (*QuestionArray, error)
}

// UnimplementedSurveyServer can be embedded to have forward compatible implementations.
type UnimplementedSurveyServer struct {
}

func (*UnimplementedSurveyServer) CreateSurvey(ctx context.Context, req *SurveyData) (*SurveyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSurvey not implemented")
}
func (*UnimplementedSurveyServer) DeleteSurvey(ctx context.Context, req *SurveyID) (*EmptySurvey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSurvey not implemented")
}
func (*UnimplementedSurveyServer) GetSurvey(ctx context.Context, req *SurveyID) (*SurveyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSurvey not implemented")
}
func (*UnimplementedSurveyServer) GetAllSurveys(ctx context.Context, req *EmptySurvey) (*SurveyArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSurveys not implemented")
}
func (*UnimplementedSurveyServer) CreateQuestion(ctx context.Context, req *Question) (*Question, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestion not implemented")
}
func (*UnimplementedSurveyServer) DeleteQuestion(ctx context.Context, req *QuestionID) (*EmptySurvey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuestion not implemented")
}
func (*UnimplementedSurveyServer) GetQuestion(ctx context.Context, req *QuestionID) (*Question, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestion not implemented")
}
func (*UnimplementedSurveyServer) GetAllQuestions(ctx context.Context, req *EmptySurvey) (*QuestionArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllQuestions not implemented")
}

func RegisterSurveyServer(s *grpc.Server, srv SurveyServer) {
	s.RegisterService(&_Survey_serviceDesc, srv)
}

func _Survey_CreateSurvey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SurveyData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurveyServer).CreateSurvey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Survey/CreateSurvey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurveyServer).CreateSurvey(ctx, req.(*SurveyData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Survey_DeleteSurvey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SurveyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurveyServer).DeleteSurvey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Survey/DeleteSurvey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurveyServer).DeleteSurvey(ctx, req.(*SurveyID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Survey_GetSurvey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SurveyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurveyServer).GetSurvey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Survey/GetSurvey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurveyServer).GetSurvey(ctx, req.(*SurveyID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Survey_GetAllSurveys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptySurvey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurveyServer).GetAllSurveys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Survey/GetAllSurveys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurveyServer).GetAllSurveys(ctx, req.(*EmptySurvey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Survey_CreateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Question)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurveyServer).CreateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Survey/CreateQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurveyServer).CreateQuestion(ctx, req.(*Question))
	}
	return interceptor(ctx, in, info, handler)
}

func _Survey_DeleteQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurveyServer).DeleteQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Survey/DeleteQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurveyServer).DeleteQuestion(ctx, req.(*QuestionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Survey_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurveyServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Survey/GetQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurveyServer).GetQuestion(ctx, req.(*QuestionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Survey_GetAllQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptySurvey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurveyServer).GetAllQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Survey/GetAllQuestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurveyServer).GetAllQuestions(ctx, req.(*EmptySurvey))
	}
	return interceptor(ctx, in, info, handler)
}

var _Survey_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Survey",
	HandlerType: (*SurveyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSurvey",
			Handler:    _Survey_CreateSurvey_Handler,
		},
		{
			MethodName: "DeleteSurvey",
			Handler:    _Survey_DeleteSurvey_Handler,
		},
		{
			MethodName: "GetSurvey",
			Handler:    _Survey_GetSurvey_Handler,
		},
		{
			MethodName: "GetAllSurveys",
			Handler:    _Survey_GetAllSurveys_Handler,
		},
		{
			MethodName: "CreateQuestion",
			Handler:    _Survey_CreateQuestion_Handler,
		},
		{
			MethodName: "DeleteQuestion",
			Handler:    _Survey_DeleteQuestion_Handler,
		},
		{
			MethodName: "GetQuestion",
			Handler:    _Survey_GetQuestion_Handler,
		},
		{
			MethodName: "GetAllQuestions",
			Handler:    _Survey_GetAllQuestions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "survey.proto",
}
