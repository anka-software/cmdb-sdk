package models

type IdentifyReconcileItem struct {
	ClassName           string               `json:"className"`
	InternalId          string               `json:"internal_id,omitempty"`
	Lookup              []*Lookup            `json:"lookup,omitempty"`
	Related             []string             `json:"related,omitempty"`
	Settings            *Settings            `json:"settings,omitempty"`
	SysObjectSourceInfo *SysObjectSourceInfo `json:"sys_object_source_info,omitempty"`
	Values              *ItemValue           `json:"values"`
	Result              *Result              `json:"result"`
}

// For response
type Items struct {
	ClassName string `json:"className"`
	Operation string `json:"operation,omitempty"`
	SysId     string `json:"sysId,omitempty"`
}
type Result struct {
	Items *[]Items `json:"items"`
}

type IdentifyReconcileItemList struct {
	Items []*IdentifyReconcileItem `json:"items"`
}

type Lookup struct {
	ClassName           string               `json:"className"`
	InternalId          string               `json:"internal_id,omitempty"`
	SysObjectSourceInfo *SysObjectSourceInfo `json:"sys_object_source_info,omitempty"`
	Values              string               `json:"values"`
}

type SysObjectSourceInfo struct {
	SourceFeed             string `json:"source_feed"`
	SourceName             string `json:"source_name"`
	SourceNativeKey        string `json:"source_native_key"`
	SourceRecencyTimestamp string `json:"source_recency_timestamp"`
}

/*type ItemValue struct {
	HostName     string `json:"host_name,omitempty"`
	IpAddress    string `json:"ip_address,omitempty"`
	Name         string `json:"name"`
	OsName       string `json:"os_name,omitempty"`
	SysClassName string `json:"sys_class_name,omitempty"`
	Ram          string `json:"ram,omitempty"`
}*/
type ItemValue struct {
	Values []map[string]interface{}
}

type Settings struct {
	SkipReclassificationRestrictionRules bool `json:"skipReclassificationRestrictionRules,omitempty"`
	UpdateWithoutDowngrade               bool `json:"updateWithoutDowngrade,omitempty"`
	UpdateWithoutSwitch                  bool `json:"updateWithoutSwitch,omitempty"`
	UpdateWithoutUpgrade                 bool `json:"updateWithoutUpgrade,omitempty"`
}

/*
// Validate validates this iaas  project
func (m *IaaSProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdministrators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IaaSProject) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	for k := range m.Links {

		if err := validate.Required("_links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}
		if val, ok := m.Links[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_links" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("_links" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) validateAdministrators(formats strfmt.Registry) error {
	if swag.IsZero(m.Administrators) { // not required
		return nil
	}

	for i := 0; i < len(m.Administrators); i++ {
		if swag.IsZero(m.Administrators[i]) { // not required
			continue
		}

		if m.Administrators[i] != nil {
			if err := m.Administrators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("administrators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("administrators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) validateConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.Constraints) { // not required
		return nil
	}

	for k := range m.Constraints {

		if err := validate.Required("constraints"+"."+k, "body", m.Constraints[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.Constraints[k]); i++ {

			if err := m.Constraints[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("constraints" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *IaaSProject) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *IaaSProject) validateMembers(formats strfmt.Registry) error {
	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) validateViewers(formats strfmt.Registry) error {
	if swag.IsZero(m.Viewers) { // not required
		return nil
	}

	for i := 0; i < len(m.Viewers); i++ {
		if swag.IsZero(m.Viewers[i]) { // not required
			continue
		}

		if m.Viewers[i] != nil {
			if err := m.Viewers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("viewers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("viewers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) validateZones(formats strfmt.Registry) error {
	if swag.IsZero(m.Zones) { // not required
		return nil
	}

	for i := 0; i < len(m.Zones); i++ {
		if swag.IsZero(m.Zones[i]) { // not required
			continue
		}

		if m.Zones[i] != nil {
			if err := m.Zones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this iaas  project based on the context it is used
func (m *IaaSProject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdministrators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMembers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateViewers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IaaSProject) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	for k := range m.Links {

		if val, ok := m.Links[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) contextValidateAdministrators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Administrators); i++ {

		if m.Administrators[i] != nil {
			if err := m.Administrators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("administrators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("administrators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) contextValidateConstraints(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Constraints {

		for i := 0; i < len(m.Constraints[k]); i++ {

			if err := m.Constraints[k][i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("constraints" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *IaaSProject) contextValidateMembers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Members); i++ {

		if m.Members[i] != nil {
			if err := m.Members[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) contextValidateViewers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Viewers); i++ {

		if m.Viewers[i] != nil {
			if err := m.Viewers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("viewers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("viewers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) contextValidateZones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Zones); i++ {

		if m.Zones[i] != nil {
			if err := m.Zones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IaaSProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IaaSProject) UnmarshalBinary(b []byte) error {
	var res IaaSProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
*/
