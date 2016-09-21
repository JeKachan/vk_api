package vk_api

type VkApiParams struct {
	METHOD_NAME string
	PARAMETERS string
	ACCESS_TOKEN string
	V string
}

func NewVkApiParams() *VkApiParams {
	model := new(VkApiParams)
	model.V = "5.51"
	return model
}

