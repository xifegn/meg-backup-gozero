// Code generated by goctl. DO NOT EDIT.
package types

type PredictRequest struct {
	Code string `form:"code"`
}

type PredictResponse struct {
	Message string `json:"message"`
}

type UploadRequest struct {
	Pid string `json:"pid"`
}

type UploadResponse struct {
	Message string `json:"message"`
}

type SmUploadResponse struct {
	Url string `json:"url"`
}

type DownloadFileRequest struct {
	Filename string `path:"filename"`
}

type DownloadFileResponse struct {
	Data []byte `json:"data"`
}

type SendPidRequest struct {
	Current_pid string `json:"currentPid"`
}

type SendPidResponse struct {
	Message string `json:"message"`
}

type GetPidResponse struct {
	Message    string `json:"message"`
	CurrentPid string `json:"currentPid"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  int64  `json:"isAdmin,omitempty,default=0"`
	Name     string `json:"name"`
	Number   string `json:"number"`
}

type RegisterResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ChangePasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ChangePasswordResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ChangePhoneNumberRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phonenumber"`
}

type ChangePhoneNumberResponse struct {
	Message string `json:"message"`
}

type ChangeEmailRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ChangeEmailResponse struct {
	Message string `json:"message"`
}

type InsertDoctorInfoRequest struct {
	Email           string `json:"email"`
	Nickname        string `json:"nickname"`
	Regions         string `json:"regions"`
	SelfInformation string `json:"selfInformation"`
	Did             string `json:"did"`
}

type InsertDoctorInfoResponse struct {
	Message string `json:"message"`
}

type GetDoctorInfoRequest struct {
	Username string `form:"username"`
}

type GetDoctorInfoResponse struct {
	Username        string `json:"username"`
	Phonenumber     string `json:"phonenumber"`
	Id              int64  `json:"id"`
	SecretKey       string `json:"secretKey"`
	Email           string `json:"email"`
	Nickname        string `json:"nickname"`
	Regions         string `json:"regions"`
	SelfInformation string `json:"selfInformation"`
}

type InsertCaseRequest struct {
	Body string `json:"body"`
	Cid  string `json:"cid"`
}

type InsertCaseResponse struct {
	Message string `json:"message"`
}

type QueryCaseRequest struct {
	Cid string `form:"cid"`
}

type QueryCaseResponse struct {
	Body string `json:"body"`
}

type CardInfo struct {
	CardValue string `json:"cardValue"`
	Name      string `json:"name"`
	Value     int64  `json:"value"`
}

type GetDataRequest struct {
	Username string `form:"username"`
}

type GetDataResponse struct {
	Data []CardInfo `json:"data"`
}

type RemoveFilePathRequest struct {
	FilePath string `json:"filePath"`
}

type RemoveFilePathResponse struct {
	Message string `json:"message"`
}

type QuotaApplyInqueryResponse struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Quota     int64  `json:"quota"`
	Amount    int64  `json:"amount"`
	CreatedAt string `json:"createdAt"`
}

type QuotaApplyRemoveRequest struct {
	Id int64 `json:"id"`
}

type QuotaApplyRemoveResponse struct {
	Message string `json:"message"`
}

type QuotaInquiryRequest struct {
	Username string `form:"username"`
}

type DailyCallInfo struct {
	DailyCallLimit string `json:"DailyCallLimit"`
	CallNumbers    string `json:"callNumbers"`
	CallVolumes    string `json:"callVolumes"`
	Type           string `json:"type"`
}

type QuotaInquiryResponse struct {
	Data []DailyCallInfo `json:"data"`
}

type GetAllRequest struct {
	Did string `form:"did"`
}

type PatientInfo struct {
	Id         int64  `json:"id"`
	Did        string `json:"did"`
	Name       string `json:"name"`
	Sex        string `json:"sex"`
	Age        int64  `json:"age"`
	UploadTime string `json:"uploadTime"`
	Code       string `json:"code"`
	FilePath   string `json:"filePath"`
}

type GetAllResponse struct {
	Data []PatientInfo `json:"data"`
}

type GetPatientFilePathRequest struct {
	Code string `json:"code"`
}

type GetPatientFilePathResponse struct {
	Key      int64  `json:"key"`
	FileName string `json:"fileName"`
}

type InsertPatientRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Age      int64  `json:"age"`
	Code     int64  `json:"code"`
}

type InsertPatientResponse struct {
	Message string `json:"message"`
}

type CardInfoa struct {
	CardValue string `json:"cardValue"`
	Name      string `json:"name"`
	Value     int64  `json:"value"`
}

type GetAllDataResponse struct {
	Data []CardInfoa `json:"data"`
}

type RemoveDoctorRequest struct {
	Username string `json:"username"`
}

type RemoveDoctorResponse struct {
	Message string `json:"message"`
}

type RemovePatientRequest struct {
	Id string `json:"id"`
}

type RemovePatientResponse struct {
	Message string `json:"message"`
}

type SeriesItem struct {
	Data []int64 `json:"data"`
	Name string  `json:"name"`
	Type string  `json:"type"`
}

type XAxis struct {
	Data []string `json:"data"`
}

type ChartDataResponse struct {
	Max    int64        `json:"max"`
	Series []SeriesItem `json:"series"`
	XAxis  XAxis        `json:"xAxis"`
}

type QuotaApplyInsertRequest struct {
	Username string `json:"username"`
	Amount   int64  `json:"amount"`
}

type QuotaApplyInsertResponse struct {
	Message string `json:"message"`
}

type AddQuotaRequest struct {
	Username string `json:"username"`
	Amount   int64  `json:"amount"`
}

type AddQuotaResponse struct {
	Message string `json:"message"`
}

type DoctorInfo struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	RegisterTime string `json:"registerTime"`
	Telephone    string `json:"telephone"`
	Username     string `json:"username"`
}

type AdminGetAllDoctorResponse struct {
	Data []DoctorInfo `json:"data"`
}

type TestResponse struct {
	Message string `json:"message"`
}
