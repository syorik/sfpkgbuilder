package pkg

const (
	ApexClassMdt                   = "ApexClass"
	ApexTriggerMdt                 = "ApexTrigger"
	CustomFieldMdt                 = "CustomField"
	CustomObjectMdt                = "CustomObject"
	CustomObjectTranslationMdt     = "CustomObjectTranslation"
	CustomPermissionMdt            = "CustomPermission"
	CustomTabMdt                   = "CustomTab"
	ExperienceBundleMdt            = "ExperienceBundle"
	FlexiPageMdt                   = "FlexiPage"
	FlowMdt                        = "Flow"
	GlobalPicklistMdt              = "GlobalPicklist"
	GlobalValueSetMdt              = "GlobalValueSet"
	GlobalValueSetTranslationMdt   = "GlobalValueSetTranslation"
	LightningComponentBundleMdt    = "LightningComponentBundle"
	LightningMessageChannelMdt     = "LightningMessageChannel"
	ListViewMdt                    = "ListView"
	MilestoneTypeMdt               = "MilestoneType"
	PermissionSetMdt               = "PermissionSet"
	PermissionSetGroupMdt          = "PermissionSetGroup"
	PlatformEventChannelMdt        = "PlatformEventChannel"
	ProfileMdt                     = "Profile"
	QueueMdt                       = "Queue"
	StandardValueSetMdt            = "StandardValueSet"
	StandardValueSetTranslationMdt = "StandardValueSetTranslation"
	StaticResourceMdt              = "StaticResource"
)

func MapDirectoryToMetadataType(directory string) string {
	switch directory {
	case "classes":
		return ApexClassMdt
	case "triggers":
		return ApexTriggerMdt
	case "objects":
		return CustomObjectMdt
	case "fields":
		return CustomFieldMdt
	case "objectTranslations":
		return CustomObjectTranslationMdt
	case "customPermissions":
		return CustomPermissionMdt
	case "tabs":
		return CustomTabMdt
	case "experiences":
		return ExperienceBundleMdt
	case "flexipages":
		return FlexiPageMdt
	case "flows":
		return FlowMdt
	case "globalPicklists":
		return GlobalPicklistMdt
	case "globalValueSets":
		return GlobalValueSetMdt
	case "globalValueSetTranslations":
		return GlobalValueSetTranslationMdt
	case "lwc":
		return LightningComponentBundleMdt
	case "messageChannels":
		return LightningMessageChannelMdt
	case "milestoneTypes":
		return MilestoneTypeMdt
	case "permissionsets":
		return PermissionSetMdt
	case "permissionsetgroups":
		return PermissionSetGroupMdt
	case "platformEventChannels":
		return PlatformEventChannelMdt
	case "profiles":
		return ProfileMdt
	case "queues":
		return QueueMdt
	case "standardValueSets":
		return StandardValueSetMdt
	case "standardValueSetTranslations":
		return StandardValueSetTranslationMdt
	case "staticresources":
		return StaticResourceMdt
	case "listViews":
		return ListViewMdt
	}
	return ""
}
