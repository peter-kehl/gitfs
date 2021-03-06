// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Basic struct {
	Username             string                `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             *wrappers.StringValue `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Basic) Reset()         { *m = Basic{} }
func (m *Basic) String() string { return proto.CompactTextString(m) }
func (*Basic) ProtoMessage()    {}
func (*Basic) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_311c42415dba4c4f, []int{0}
}
func (m *Basic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Basic.Unmarshal(m, b)
}
func (m *Basic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Basic.Marshal(b, m, deterministic)
}
func (dst *Basic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Basic.Merge(dst, src)
}
func (m *Basic) XXX_Size() int {
	return xxx_messageInfo_Basic.Size(m)
}
func (m *Basic) XXX_DiscardUnknown() {
	xxx_messageInfo_Basic.DiscardUnknown(m)
}

var xxx_messageInfo_Basic proto.InternalMessageInfo

func (m *Basic) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Basic) GetPassword() *wrappers.StringValue {
	if m != nil {
		return m.Password
	}
	return nil
}

type OAuthToken struct {
	Token                string                `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ApplicationId        *wrappers.StringValue `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OAuthToken) Reset()         { *m = OAuthToken{} }
func (m *OAuthToken) String() string { return proto.CompactTextString(m) }
func (*OAuthToken) ProtoMessage()    {}
func (*OAuthToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_311c42415dba4c4f, []int{1}
}
func (m *OAuthToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthToken.Unmarshal(m, b)
}
func (m *OAuthToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthToken.Marshal(b, m, deterministic)
}
func (dst *OAuthToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthToken.Merge(dst, src)
}
func (m *OAuthToken) XXX_Size() int {
	return xxx_messageInfo_OAuthToken.Size(m)
}
func (m *OAuthToken) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthToken.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthToken proto.InternalMessageInfo

func (m *OAuthToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *OAuthToken) GetApplicationId() *wrappers.StringValue {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

type OAuth2Token struct {
	Token                string                `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	TokenType            *wrappers.StringValue `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	RefreshToken         *wrappers.StringValue `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Expiry               *wrappers.StringValue `protobuf:"bytes,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OAuth2Token) Reset()         { *m = OAuth2Token{} }
func (m *OAuth2Token) String() string { return proto.CompactTextString(m) }
func (*OAuth2Token) ProtoMessage()    {}
func (*OAuth2Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_311c42415dba4c4f, []int{2}
}
func (m *OAuth2Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuth2Token.Unmarshal(m, b)
}
func (m *OAuth2Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuth2Token.Marshal(b, m, deterministic)
}
func (dst *OAuth2Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuth2Token.Merge(dst, src)
}
func (m *OAuth2Token) XXX_Size() int {
	return xxx_messageInfo_OAuth2Token.Size(m)
}
func (m *OAuth2Token) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuth2Token.DiscardUnknown(m)
}

var xxx_messageInfo_OAuth2Token proto.InternalMessageInfo

func (m *OAuth2Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *OAuth2Token) GetTokenType() *wrappers.StringValue {
	if m != nil {
		return m.TokenType
	}
	return nil
}

func (m *OAuth2Token) GetRefreshToken() *wrappers.StringValue {
	if m != nil {
		return m.RefreshToken
	}
	return nil
}

func (m *OAuth2Token) GetExpiry() *wrappers.StringValue {
	if m != nil {
		return m.Expiry
	}
	return nil
}

type Github struct {
	BaseUrl              *wrappers.StringValue `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	UploadUrl            *wrappers.StringValue `protobuf:"bytes,2,opt,name=upload_url,json=uploadUrl,proto3" json:"upload_url,omitempty"`
	Users                []string              `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	Organizations        []string              `protobuf:"bytes,4,rep,name=organizations,proto3" json:"organizations,omitempty"`
	Oauth2               *OAuth2Token          `protobuf:"bytes,10,opt,name=oauth2,proto3" json:"oauth2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Github) Reset()         { *m = Github{} }
func (m *Github) String() string { return proto.CompactTextString(m) }
func (*Github) ProtoMessage()    {}
func (*Github) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_311c42415dba4c4f, []int{3}
}
func (m *Github) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Github.Unmarshal(m, b)
}
func (m *Github) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Github.Marshal(b, m, deterministic)
}
func (dst *Github) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Github.Merge(dst, src)
}
func (m *Github) XXX_Size() int {
	return xxx_messageInfo_Github.Size(m)
}
func (m *Github) XXX_DiscardUnknown() {
	xxx_messageInfo_Github.DiscardUnknown(m)
}

var xxx_messageInfo_Github proto.InternalMessageInfo

func (m *Github) GetBaseUrl() *wrappers.StringValue {
	if m != nil {
		return m.BaseUrl
	}
	return nil
}

func (m *Github) GetUploadUrl() *wrappers.StringValue {
	if m != nil {
		return m.UploadUrl
	}
	return nil
}

func (m *Github) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Github) GetOrganizations() []string {
	if m != nil {
		return m.Organizations
	}
	return nil
}

func (m *Github) GetOauth2() *OAuth2Token {
	if m != nil {
		return m.Oauth2
	}
	return nil
}

type Gitlab struct {
	BaseUrl              *wrappers.StringValue `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	Private              *OAuthToken           `protobuf:"bytes,10,opt,name=private,proto3" json:"private,omitempty"`
	Oauth                *OAuthToken           `protobuf:"bytes,11,opt,name=oauth,proto3" json:"oauth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Gitlab) Reset()         { *m = Gitlab{} }
func (m *Gitlab) String() string { return proto.CompactTextString(m) }
func (*Gitlab) ProtoMessage()    {}
func (*Gitlab) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_311c42415dba4c4f, []int{4}
}
func (m *Gitlab) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gitlab.Unmarshal(m, b)
}
func (m *Gitlab) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gitlab.Marshal(b, m, deterministic)
}
func (dst *Gitlab) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gitlab.Merge(dst, src)
}
func (m *Gitlab) XXX_Size() int {
	return xxx_messageInfo_Gitlab.Size(m)
}
func (m *Gitlab) XXX_DiscardUnknown() {
	xxx_messageInfo_Gitlab.DiscardUnknown(m)
}

var xxx_messageInfo_Gitlab proto.InternalMessageInfo

func (m *Gitlab) GetBaseUrl() *wrappers.StringValue {
	if m != nil {
		return m.BaseUrl
	}
	return nil
}

func (m *Gitlab) GetPrivate() *OAuthToken {
	if m != nil {
		return m.Private
	}
	return nil
}

func (m *Gitlab) GetOauth() *OAuthToken {
	if m != nil {
		return m.Oauth
	}
	return nil
}

type Bitbucket struct {
	Basic                *Basic      `protobuf:"bytes,10,opt,name=basic,proto3" json:"basic,omitempty"`
	Oauth                *OAuthToken `protobuf:"bytes,11,opt,name=oauth,proto3" json:"oauth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Bitbucket) Reset()         { *m = Bitbucket{} }
func (m *Bitbucket) String() string { return proto.CompactTextString(m) }
func (*Bitbucket) ProtoMessage()    {}
func (*Bitbucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_311c42415dba4c4f, []int{5}
}
func (m *Bitbucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bitbucket.Unmarshal(m, b)
}
func (m *Bitbucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bitbucket.Marshal(b, m, deterministic)
}
func (dst *Bitbucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bitbucket.Merge(dst, src)
}
func (m *Bitbucket) XXX_Size() int {
	return xxx_messageInfo_Bitbucket.Size(m)
}
func (m *Bitbucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bitbucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bitbucket proto.InternalMessageInfo

func (m *Bitbucket) GetBasic() *Basic {
	if m != nil {
		return m.Basic
	}
	return nil
}

func (m *Bitbucket) GetOauth() *OAuthToken {
	if m != nil {
		return m.Oauth
	}
	return nil
}

type Generic struct {
	BaseUrl              string   `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	PerPageParameter     string   `protobuf:"bytes,3,opt,name=per_page_parameter,json=perPageParameter,proto3" json:"per_page_parameter,omitempty"`
	PageParameter        string   `protobuf:"bytes,4,opt,name=page_parameter,json=pageParameter,proto3" json:"page_parameter,omitempty"`
	PageSize             int32    `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Selector             string   `protobuf:"bytes,6,opt,name=selector,proto3" json:"selector,omitempty"`
	Basic                *Basic   `protobuf:"bytes,10,opt,name=basic,proto3" json:"basic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Generic) Reset()         { *m = Generic{} }
func (m *Generic) String() string { return proto.CompactTextString(m) }
func (*Generic) ProtoMessage()    {}
func (*Generic) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_311c42415dba4c4f, []int{6}
}
func (m *Generic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Generic.Unmarshal(m, b)
}
func (m *Generic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Generic.Marshal(b, m, deterministic)
}
func (dst *Generic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Generic.Merge(dst, src)
}
func (m *Generic) XXX_Size() int {
	return xxx_messageInfo_Generic.Size(m)
}
func (m *Generic) XXX_DiscardUnknown() {
	xxx_messageInfo_Generic.DiscardUnknown(m)
}

var xxx_messageInfo_Generic proto.InternalMessageInfo

func (m *Generic) GetBaseUrl() string {
	if m != nil {
		return m.BaseUrl
	}
	return ""
}

func (m *Generic) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Generic) GetPerPageParameter() string {
	if m != nil {
		return m.PerPageParameter
	}
	return ""
}

func (m *Generic) GetPageParameter() string {
	if m != nil {
		return m.PageParameter
	}
	return ""
}

func (m *Generic) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Generic) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *Generic) GetBasic() *Basic {
	if m != nil {
		return m.Basic
	}
	return nil
}

type Account struct {
	Github               *Github    `protobuf:"bytes,1,opt,name=github,proto3" json:"github,omitempty"`
	Gitlab               *Gitlab    `protobuf:"bytes,2,opt,name=gitlab,proto3" json:"gitlab,omitempty"`
	Bitbucket            *Bitbucket `protobuf:"bytes,3,opt,name=bitbucket,proto3" json:"bitbucket,omitempty"`
	Generic              *Generic   `protobuf:"bytes,4,opt,name=generic,proto3" json:"generic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_311c42415dba4c4f, []int{7}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (dst *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(dst, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetGithub() *Github {
	if m != nil {
		return m.Github
	}
	return nil
}

func (m *Account) GetGitlab() *Gitlab {
	if m != nil {
		return m.Gitlab
	}
	return nil
}

func (m *Account) GetBitbucket() *Bitbucket {
	if m != nil {
		return m.Bitbucket
	}
	return nil
}

func (m *Account) GetGeneric() *Generic {
	if m != nil {
		return m.Generic
	}
	return nil
}

type CloneOverride struct {
	RepositoryRoot       *wrappers.StringValue `protobuf:"bytes,1,opt,name=repository_root,json=repositoryRoot,proto3" json:"repository_root,omitempty"`
	Depth                *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=depth,proto3" json:"depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CloneOverride) Reset()         { *m = CloneOverride{} }
func (m *CloneOverride) String() string { return proto.CompactTextString(m) }
func (*CloneOverride) ProtoMessage()    {}
func (*CloneOverride) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_311c42415dba4c4f, []int{8}
}
func (m *CloneOverride) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloneOverride.Unmarshal(m, b)
}
func (m *CloneOverride) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloneOverride.Marshal(b, m, deterministic)
}
func (dst *CloneOverride) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloneOverride.Merge(dst, src)
}
func (m *CloneOverride) XXX_Size() int {
	return xxx_messageInfo_CloneOverride.Size(m)
}
func (m *CloneOverride) XXX_DiscardUnknown() {
	xxx_messageInfo_CloneOverride.DiscardUnknown(m)
}

var xxx_messageInfo_CloneOverride proto.InternalMessageInfo

func (m *CloneOverride) GetRepositoryRoot() *wrappers.StringValue {
	if m != nil {
		return m.RepositoryRoot
	}
	return nil
}

func (m *CloneOverride) GetDepth() *wrappers.Int32Value {
	if m != nil {
		return m.Depth
	}
	return nil
}

type CloneConfiguration struct {
	RepositoryRoot       *wrappers.StringValue     `protobuf:"bytes,1,opt,name=repository_root,json=repositoryRoot,proto3" json:"repository_root,omitempty"`
	Depth                *wrappers.Int32Value      `protobuf:"bytes,2,opt,name=depth,proto3" json:"depth,omitempty"`
	Overrides            map[string]*CloneOverride `protobuf:"bytes,10,rep,name=overrides,proto3" json:"overrides,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CloneConfiguration) Reset()         { *m = CloneConfiguration{} }
func (m *CloneConfiguration) String() string { return proto.CompactTextString(m) }
func (*CloneConfiguration) ProtoMessage()    {}
func (*CloneConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_311c42415dba4c4f, []int{9}
}
func (m *CloneConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloneConfiguration.Unmarshal(m, b)
}
func (m *CloneConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloneConfiguration.Marshal(b, m, deterministic)
}
func (dst *CloneConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloneConfiguration.Merge(dst, src)
}
func (m *CloneConfiguration) XXX_Size() int {
	return xxx_messageInfo_CloneConfiguration.Size(m)
}
func (m *CloneConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_CloneConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_CloneConfiguration proto.InternalMessageInfo

func (m *CloneConfiguration) GetRepositoryRoot() *wrappers.StringValue {
	if m != nil {
		return m.RepositoryRoot
	}
	return nil
}

func (m *CloneConfiguration) GetDepth() *wrappers.Int32Value {
	if m != nil {
		return m.Depth
	}
	return nil
}

func (m *CloneConfiguration) GetOverrides() map[string]*CloneOverride {
	if m != nil {
		return m.Overrides
	}
	return nil
}

type Configuration struct {
	Mount                string              `protobuf:"bytes,1,opt,name=mount,proto3" json:"mount,omitempty"`
	Accounts             []*Account          `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	Clone                *CloneConfiguration `protobuf:"bytes,3,opt,name=clone,proto3" json:"clone,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_311c42415dba4c4f, []int{10}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (dst *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(dst, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetMount() string {
	if m != nil {
		return m.Mount
	}
	return ""
}

func (m *Configuration) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *Configuration) GetClone() *CloneConfiguration {
	if m != nil {
		return m.Clone
	}
	return nil
}

func init() {
	proto.RegisterType((*Basic)(nil), "mjpitz.gitfs.Basic")
	proto.RegisterType((*OAuthToken)(nil), "mjpitz.gitfs.OAuthToken")
	proto.RegisterType((*OAuth2Token)(nil), "mjpitz.gitfs.OAuth2Token")
	proto.RegisterType((*Github)(nil), "mjpitz.gitfs.Github")
	proto.RegisterType((*Gitlab)(nil), "mjpitz.gitfs.Gitlab")
	proto.RegisterType((*Bitbucket)(nil), "mjpitz.gitfs.Bitbucket")
	proto.RegisterType((*Generic)(nil), "mjpitz.gitfs.Generic")
	proto.RegisterType((*Account)(nil), "mjpitz.gitfs.Account")
	proto.RegisterType((*CloneOverride)(nil), "mjpitz.gitfs.CloneOverride")
	proto.RegisterType((*CloneConfiguration)(nil), "mjpitz.gitfs.CloneConfiguration")
	proto.RegisterMapType((map[string]*CloneOverride)(nil), "mjpitz.gitfs.CloneConfiguration.OverridesEntry")
	proto.RegisterType((*Configuration)(nil), "mjpitz.gitfs.Configuration")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_config_311c42415dba4c4f) }

var fileDescriptor_config_311c42415dba4c4f = []byte{
	// 811 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x96, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0xe5, 0xa6, 0x4e, 0xe2, 0x93, 0xa6, 0xac, 0x86, 0x22, 0xbc, 0x5d, 0x2e, 0x22, 0x0b,
	0xa4, 0x22, 0xad, 0x1c, 0x35, 0xcb, 0xc7, 0x02, 0x57, 0x6d, 0xb5, 0x5a, 0xed, 0x05, 0xda, 0x95,
	0x77, 0x41, 0x02, 0x09, 0x45, 0x63, 0xe7, 0xc4, 0x1d, 0xea, 0x78, 0x46, 0x33, 0xe3, 0x2e, 0xe9,
	0x1b, 0x70, 0xc7, 0x4b, 0xf0, 0x0e, 0xbc, 0x04, 0x0f, 0xc0, 0x43, 0x70, 0xc1, 0x1b, 0xa0, 0xf9,
	0x70, 0xd3, 0xa4, 0x88, 0xa6, 0xe2, 0x82, 0xbb, 0x99, 0xcc, 0xef, 0xcc, 0xff, 0x7c, 0x8e, 0x03,
	0x7b, 0x05, 0xaf, 0xe7, 0xac, 0x4c, 0x85, 0xe4, 0x9a, 0x93, 0xbd, 0xc5, 0x8f, 0x82, 0xe9, 0xab,
	0xb4, 0x64, 0x7a, 0xae, 0x0e, 0xbf, 0x28, 0x99, 0x3e, 0x6f, 0xf2, 0xb4, 0xe0, 0x8b, 0x71, 0xc9,
	0x2b, 0x5a, 0x97, 0x63, 0x8b, 0xe5, 0xcd, 0x7c, 0x2c, 0xf4, 0x52, 0xa0, 0x1a, 0xbf, 0x95, 0x54,
	0x08, 0x94, 0xab, 0x85, 0xbb, 0x28, 0xf9, 0x01, 0xc2, 0x53, 0xaa, 0x58, 0x41, 0x0e, 0xa1, 0xdf,
	0x28, 0x94, 0x35, 0x5d, 0x60, 0x1c, 0x8c, 0x82, 0xa3, 0x28, 0xbb, 0xde, 0x93, 0xa7, 0xd0, 0x17,
	0x54, 0xa9, 0xb7, 0x5c, 0xce, 0xe2, 0x9d, 0x51, 0x70, 0x34, 0x98, 0x7c, 0x90, 0x96, 0x9c, 0x97,
	0x15, 0xa6, 0xad, 0x4e, 0xfa, 0x5a, 0x4b, 0x56, 0x97, 0xdf, 0xd2, 0xaa, 0xc1, 0xec, 0x9a, 0x4e,
	0x4a, 0x80, 0x97, 0x27, 0x8d, 0x3e, 0x7f, 0xc3, 0x2f, 0xb0, 0x26, 0x07, 0x10, 0x6a, 0xb3, 0xf0,
	0x02, 0x6e, 0x43, 0xce, 0x60, 0x9f, 0x0a, 0x51, 0xb1, 0x82, 0x6a, 0xc6, 0xeb, 0x29, 0xdb, 0x4e,
	0x63, 0x78, 0xc3, 0xe6, 0xc5, 0x2c, 0xf9, 0x23, 0x80, 0x81, 0x55, 0x9a, 0xfc, 0x9b, 0xd4, 0x57,
	0x00, 0x76, 0x31, 0x35, 0x49, 0xd9, 0x4a, 0x26, 0xb2, 0xfc, 0x9b, 0xa5, 0x40, 0x72, 0x02, 0x43,
	0x89, 0x73, 0x89, 0xea, 0x7c, 0xea, 0xae, 0xee, 0x6c, 0x61, 0xbf, 0xe7, 0x4d, 0x9c, 0x57, 0x9f,
	0x40, 0x17, 0x7f, 0x12, 0x4c, 0x2e, 0xe3, 0xdd, 0x2d, 0x6c, 0x3d, 0x9b, 0xfc, 0x19, 0x40, 0xf7,
	0xb9, 0xad, 0x30, 0xf9, 0x1c, 0xfa, 0x39, 0x55, 0x38, 0x6d, 0x64, 0x65, 0x23, 0xbb, 0xeb, 0x8a,
	0x9e, 0xa1, 0xbf, 0x91, 0x95, 0x89, 0xbc, 0x11, 0x15, 0xa7, 0x33, 0x6b, 0xba, 0x55, 0xe4, 0x8e,
	0x37, 0xc6, 0x07, 0x10, 0x9a, 0x5e, 0x50, 0x71, 0x67, 0xd4, 0x31, 0xc9, 0xb4, 0x1b, 0xf2, 0x21,
	0x0c, 0xb9, 0x2c, 0x69, 0xcd, 0xae, 0x6c, 0x11, 0x54, 0xbc, 0x6b, 0x4f, 0xd7, 0x7f, 0x24, 0xc7,
	0xd0, 0xe5, 0xd4, 0xd4, 0x25, 0x06, 0x2b, 0xfa, 0x30, 0xbd, 0xd9, 0xba, 0xe9, 0x8d, 0x9a, 0x65,
	0x1e, 0x4c, 0x7e, 0x75, 0xf1, 0x56, 0xf4, 0x3f, 0xc4, 0x3b, 0x81, 0x9e, 0x90, 0xec, 0x92, 0x6a,
	0xf4, 0xba, 0xf1, 0x3f, 0xe8, 0x3a, 0xd9, 0x16, 0x24, 0x29, 0x84, 0xd6, 0x83, 0x78, 0x70, 0x87,
	0x85, 0xc3, 0x92, 0x39, 0x44, 0xa7, 0x4c, 0xe7, 0x4d, 0x71, 0x81, 0x9a, 0x7c, 0x0c, 0x61, 0x6e,
	0x06, 0xc9, 0xcb, 0xbd, 0xbb, 0x6e, 0x6c, 0x67, 0x2c, 0x73, 0xc4, 0xbd, 0x75, 0xfe, 0x0a, 0xa0,
	0xf7, 0x1c, 0x6b, 0x94, 0xac, 0x20, 0x0f, 0x37, 0x12, 0x12, 0xad, 0x42, 0x26, 0xb0, 0x2b, 0xa8,
	0x3e, 0xb7, 0xc5, 0x8d, 0x32, 0xbb, 0x26, 0x8f, 0x81, 0x08, 0x94, 0x53, 0x41, 0x4b, 0x9c, 0x0a,
	0x2a, 0xe9, 0x02, 0x35, 0x4a, 0xdb, 0xb8, 0x51, 0xf6, 0x40, 0xa0, 0x7c, 0x45, 0x4b, 0x7c, 0xd5,
	0xfe, 0x4e, 0x3e, 0x82, 0xfd, 0x0d, 0x72, 0xd7, 0x92, 0x43, 0xb1, 0x86, 0x3d, 0x82, 0xc8, 0x62,
	0x8a, 0x5d, 0x61, 0x1c, 0x8e, 0x82, 0xa3, 0xd0, 0x4c, 0x7c, 0x89, 0xaf, 0xd9, 0x15, 0x9a, 0x77,
	0x44, 0x61, 0x85, 0x85, 0xe6, 0x32, 0xee, 0xba, 0x77, 0xa4, 0xdd, 0xdf, 0x23, 0x47, 0xc9, 0xef,
	0x01, 0xf4, 0x4e, 0x8a, 0x82, 0x37, 0xb5, 0x26, 0x8f, 0xa1, 0xeb, 0x1e, 0x38, 0xdf, 0x02, 0x07,
	0xeb, 0x76, 0x6e, 0x34, 0x32, 0xcf, 0x78, 0xba, 0xa2, 0xb9, 0xef, 0xf2, 0xdb, 0x74, 0x45, 0x1d,
	0x6d, 0x1a, 0xec, 0x53, 0x88, 0xf2, 0xb6, 0x86, 0x7e, 0xa0, 0xdf, 0xdf, 0x70, 0xab, 0x3d, 0xce,
	0x56, 0x24, 0x19, 0x43, 0xaf, 0x74, 0x15, 0xf1, 0x93, 0xfc, 0xde, 0x86, 0x8a, 0x3b, 0xcc, 0x5a,
	0x2a, 0xf9, 0x39, 0x80, 0xe1, 0x59, 0xc5, 0x6b, 0x7c, 0x79, 0x89, 0x52, 0xb2, 0x19, 0x92, 0x67,
	0xf0, 0x8e, 0x44, 0xc1, 0x15, 0xd3, 0x5c, 0x2e, 0xa7, 0x92, 0x73, 0xbd, 0x55, 0x87, 0xef, 0xaf,
	0x8c, 0x32, 0xce, 0x35, 0x39, 0x86, 0x70, 0x86, 0xc2, 0x97, 0x7d, 0x30, 0x79, 0x74, 0xcb, 0xf8,
	0x45, 0xad, 0x9f, 0x4c, 0x9c, 0xad, 0x23, 0x93, 0xdf, 0x76, 0x80, 0x58, 0x5f, 0xce, 0xec, 0x27,
	0xa5, 0x91, 0x76, 0x54, 0xff, 0x3f, 0x87, 0xc8, 0xd7, 0x10, 0x71, 0x9f, 0x16, 0x15, 0xc3, 0xa8,
	0x73, 0x34, 0x98, 0x8c, 0xd7, 0xf3, 0x79, 0xdb, 0xdd, 0xb4, 0x4d, 0xa4, 0x7a, 0x56, 0x6b, 0xb9,
	0xcc, 0x56, 0x37, 0x1c, 0x7e, 0x07, 0xfb, 0xeb, 0x87, 0xe4, 0x01, 0x74, 0x2e, 0x70, 0xe9, 0x07,
	0xc6, 0x2c, 0x8d, 0x97, 0x97, 0xc6, 0x85, 0x6b, 0x2f, 0x6f, 0xcb, 0xb5, 0x77, 0x64, 0x8e, 0xfc,
	0x72, 0xe7, 0x69, 0x90, 0xfc, 0x62, 0xca, 0xb8, 0x96, 0xb5, 0x03, 0x08, 0x17, 0xa6, 0x4b, 0xdb,
	0x0f, 0x8d, 0xdd, 0x90, 0x63, 0xe8, 0x53, 0xd7, 0xbd, 0x2a, 0xde, 0xb1, 0x01, 0x6d, 0x34, 0x88,
	0xef, 0xed, 0xec, 0x1a, 0x23, 0x9f, 0x41, 0x58, 0x18, 0x59, 0xdf, 0x85, 0xa3, 0xbb, 0x12, 0x90,
	0x39, 0xfc, 0xb4, 0xff, 0x7d, 0xd7, 0xfd, 0x35, 0xc8, 0xbb, 0x36, 0xc5, 0x4f, 0xfe, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x16, 0x6f, 0xd8, 0xa8, 0x2b, 0x08, 0x00, 0x00,
}
