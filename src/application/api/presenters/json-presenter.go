package presenters

type JsonModel struct {
	Payload any `json:"payload"`
}

func Envelope(payload any) JsonModel {

	return JsonModel{
		Payload: payload,
	}

}
