// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
)

func numericFormatConfigurationSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"currency_display_format_configuration": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CurrencyDisplayFormatConfiguration.html
					Type:     schema.TypeList,
					MinItems: 1,
					MaxItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"decimal_places_configuration":    decimalPlacesConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DecimalPlacesConfiguration.html
							"negative_value_configuration":    negativeValueConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NegativeValueConfiguration.html
							"null_value_format_configuration": nullValueConfigurationSchema(),     // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NullValueFormatConfiguration.html
							"number_scale":                    stringSchema(false, enum.Validate[types.NumberScale]()),
							"prefix":                          stringSchema(false, validation.ToDiagFunc(validation.StringLenBetween(1, 128))),
							"separator_configuration":         separatorConfigurationSchema(),
							"suffix":                          stringSchema(false, validation.ToDiagFunc(validation.StringLenBetween(1, 128))),
							"symbol":                          stringSchema(false, validation.ToDiagFunc(validation.StringMatch(regexache.MustCompile(`[A-Z]{3}`), "must be a 3 character currency symbol"))),
						},
					},
				},
				"number_display_format_configuration":     numberDisplayFormatConfigurationSchema(),     // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NumberDisplayFormatConfiguration.html
				"percentage_display_format_configuration": percentageDisplayFormatConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_PercentageDisplayFormatConfiguration.html
			},
		},
	}
}

func dateTimeFormatConfigurationSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"date_time_format":                stringSchema(false, validation.ToDiagFunc(validation.StringLenBetween(1, 128))),
				"null_value_format_configuration": nullValueConfigurationSchema(),     // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NullValueFormatConfiguration.html
				"numeric_format_configuration":    numericFormatConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NumericFormatConfiguration.html
			},
		},
	}
}

func numberDisplayFormatConfigurationSchema() *schema.Schema {
	return &schema.Schema{ // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NumberDisplayFormatConfiguration.html
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"decimal_places_configuration":    decimalPlacesConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DecimalPlacesConfiguration.html
				"negative_value_configuration":    negativeValueConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NegativeValueConfiguration.html
				"null_value_format_configuration": nullValueConfigurationSchema(),     // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NullValueFormatConfiguration.html
				"number_scale":                    stringSchema(false, enum.Validate[types.NumberScale]()),
				"prefix":                          stringSchema(false, validation.ToDiagFunc(validation.StringLenBetween(1, 128))),
				"separator_configuration":         separatorConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NumericSeparatorConfiguration.html
				"suffix":                          stringSchema(false, validation.ToDiagFunc(validation.StringLenBetween(1, 128))),
			},
		},
	}
}

func percentageDisplayFormatConfigurationSchema() *schema.Schema {
	return &schema.Schema{ // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_PercentageDisplayFormatConfiguration.html
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"decimal_places_configuration":    decimalPlacesConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DecimalPlacesConfiguration.html
				"negative_value_configuration":    negativeValueConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NegativeValueConfiguration.html
				"null_value_format_configuration": nullValueConfigurationSchema(),     // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NullValueFormatConfiguration.html
				"prefix":                          stringSchema(false, validation.ToDiagFunc(validation.StringLenBetween(1, 128))),
				"separator_configuration":         separatorConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NumericSeparatorConfiguration.html
				"suffix":                          stringSchema(false, validation.ToDiagFunc(validation.StringLenBetween(1, 128))),
			},
		},
	}
}

func numberFormatConfigurationSchema() *schema.Schema {
	return &schema.Schema{ // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NumberFormatConfiguration.html
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"numeric_format_configuration": numericFormatConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NumericFormatConfiguration.html
			},
		},
	}
}

func stringFormatConfigurationSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"null_value_format_configuration": nullValueConfigurationSchema(),     // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NullValueFormatConfiguration.html
				"numeric_format_configuration":    numericFormatConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NumericFormatConfiguration.html
			},
		},
	}
}

func decimalPlacesConfigurationSchema() *schema.Schema {
	return &schema.Schema{ // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DecimalPlacesConfiguration.html
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"decimal_places": {
					Type:             schema.TypeInt,
					Required:         true,
					ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 20)),
				},
			},
		},
	}
}

func negativeValueConfigurationSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"display_mode": stringSchema(true, enum.Validate[types.NegativeValueDisplayMode]()),
			},
		},
	}
}

func nullValueConfigurationSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"null_string": stringSchema(true, validation.ToDiagFunc(validation.StringLenBetween(1, 128))),
			},
		},
	}
}

func separatorConfigurationSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"decimal_separator": stringSchema(false, enum.Validate[types.NumericSeparatorSymbol]()),
				"thousands_separator": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ThousandSeparatorOptions.html
					Type:     schema.TypeList,
					MinItems: 1,
					MaxItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"symbol":     stringSchema(false, enum.Validate[types.NumericSeparatorSymbol]()),
							"visibility": stringSchema(false, enum.Validate[types.Visibility]()),
						},
					},
				},
			},
		},
	}
}

func labelOptionsSchema() *schema.Schema {
	return &schema.Schema{ // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_LabelOptions.html
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"custom_label": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"font_configuration": fontConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_FontConfiguration.html
				"visibility":         stringSchema(false, enum.Validate[types.Visibility]()),
			},
		},
	}
}

func fontConfigurationSchema() *schema.Schema {
	return &schema.Schema{ // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_FontConfiguration.html
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"font_color":      stringSchema(false, validation.ToDiagFunc(validation.StringMatch(regexache.MustCompile(`^#[0-9A-F]{6}$`), ""))),
				"font_decoration": stringSchema(false, enum.Validate[types.FontDecoration]()),
				"font_size": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_FontSize.html
					Type:     schema.TypeList,
					MaxItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"relative": stringSchema(false, enum.Validate[types.RelativeFontSize]()),
						},
					},
				},
				"font_style": stringSchema(false, enum.Validate[types.FontStyle]()),
				"font_weight": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_FontWeight.html
					Type:     schema.TypeList,
					MaxItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": stringSchema(false, enum.Validate[types.FontWeightName]()),
						},
					},
				},
			},
		},
	}
}

func formatConfigurationSchema() *schema.Schema {
	return &schema.Schema{ // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_FormatConfiguration.html
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"date_time_format_configuration": dateTimeFormatConfigurationSchema(), // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DateTimeFormatConfiguration.html
				"number_format_configuration":    numberFormatConfigurationSchema(),   // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_NumberFormatConfiguration.html
				"string_format_configuration":    stringFormatConfigurationSchema(),   // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_StringFormatConfiguration.html
			},
		},
	}
}

func expandFormatConfiguration(tfList []interface{}) *types.FormatConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.FormatConfiguration{}

	if v, ok := tfMap["date_time_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.DateTimeFormatConfiguration = expandDateTimeFormatConfiguration(v)
	}
	if v, ok := tfMap["number_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NumberFormatConfiguration = expandNumberFormatConfiguration(v)
	}
	if v, ok := tfMap["string_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.StringFormatConfiguration = expandStringFormatConfiguration(v)
	}

	return config
}

func expandDateTimeFormatConfiguration(tfList []interface{}) *types.DateTimeFormatConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.DateTimeFormatConfiguration{}

	if v, ok := tfMap["date_time_format"].(string); ok && v != "" {
		config.DateTimeFormat = aws.String(v)
	}
	if v, ok := tfMap["null_value_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NullValueFormatConfiguration = expandNullValueFormatConfiguration(v)
	}
	if v, ok := tfMap["numeric_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NumericFormatConfiguration = expandNumericFormatConfiguration(v)
	}

	return config
}

func expandNullValueFormatConfiguration(tfList []interface{}) *types.NullValueFormatConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.NullValueFormatConfiguration{}

	if v, ok := tfMap["null_string"].(string); ok && v != "" {
		config.NullString = aws.String(v)
	}

	return config
}

func expandNumericFormatConfiguration(tfList []interface{}) *types.NumericFormatConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.NumericFormatConfiguration{}

	if v, ok := tfMap["currency_display_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.CurrencyDisplayFormatConfiguration = expandCurrencyDisplayFormatConfiguration(v)
	}
	if v, ok := tfMap["number_display_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NumberDisplayFormatConfiguration = expandNumberDisplayFormatConfiguration(v)
	}
	if v, ok := tfMap["percentage_display_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.PercentageDisplayFormatConfiguration = expandPercentageDisplayFormatConfiguration(v)
	}

	return config
}

func expandCurrencyDisplayFormatConfiguration(tfList []interface{}) *types.CurrencyDisplayFormatConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.CurrencyDisplayFormatConfiguration{}

	if v, ok := tfMap["decimal_places_configuration"].([]interface{}); ok && len(v) > 0 {
		config.DecimalPlacesConfiguration = expandDecimalPlacesConfiguration(v)
	}
	if v, ok := tfMap["negative_value_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NegativeValueConfiguration = expandNegativeValueConfiguration(v)
	}
	if v, ok := tfMap["null_value_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NullValueFormatConfiguration = expandNullValueFormatConfiguration(v)
	}
	if v, ok := tfMap["number_scale"].(string); ok && v != "" {
		config.NumberScale = types.NumberScale(v)
	}
	if v, ok := tfMap["prefix"].(string); ok && v != "" {
		config.Prefix = aws.String(v)
	}
	if v, ok := tfMap["separator_configuration"].([]interface{}); ok && len(v) > 0 {
		config.SeparatorConfiguration = expandNumericSeparatorConfiguration(v)
	}
	if v, ok := tfMap["suffix"].(string); ok && v != "" {
		config.Suffix = aws.String(v)
	}
	if v, ok := tfMap["symbol"].(string); ok && v != "" {
		config.Symbol = aws.String(v)
	}

	return config
}

func expandDecimalPlacesConfiguration(tfList []interface{}) *types.DecimalPlacesConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.DecimalPlacesConfiguration{}

	if v, ok := tfMap["decimal_places"].(int); ok {
		config.DecimalPlaces = aws.Int64(int64(v))
	}

	return config
}

func expandNegativeValueConfiguration(tfList []interface{}) *types.NegativeValueConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.NegativeValueConfiguration{}

	if v, ok := tfMap["display_mode"].(string); ok {
		config.DisplayMode = types.NegativeValueDisplayMode(v)
	}

	return config
}

func expandNumericSeparatorConfiguration(tfList []interface{}) *types.NumericSeparatorConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.NumericSeparatorConfiguration{}

	if v, ok := tfMap["decimal_separator"].(string); ok {
		config.DecimalSeparator = types.NumericSeparatorSymbol(v)
	}
	if v, ok := tfMap["thousands_separator"].([]interface{}); ok && len(v) > 0 {
		config.ThousandsSeparator = expandThousandSeparatorOptions(v)
	}

	return config
}

func expandThousandSeparatorOptions(tfList []interface{}) *types.ThousandSeparatorOptions {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.ThousandSeparatorOptions{}

	if v, ok := tfMap["symbol"].(string); ok {
		config.Symbol = types.NumericSeparatorSymbol(v)
	}
	if v, ok := tfMap["visibility"].(string); ok {
		config.Visibility = types.Visibility(v)
	}

	return config
}

func expandNumberFormatConfiguration(tfList []interface{}) *types.NumberFormatConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.NumberFormatConfiguration{}

	if v, ok := tfMap["numeric_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.FormatConfiguration = expandNumericFormatConfiguration(v)
	}

	return config
}

func expandStringFormatConfiguration(tfList []interface{}) *types.StringFormatConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.StringFormatConfiguration{}

	if v, ok := tfMap["null_value_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NullValueFormatConfiguration = expandNullValueFormatConfiguration(v)
	}
	if v, ok := tfMap["numeric_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NumericFormatConfiguration = expandNumericFormatConfiguration(v)
	}

	return config
}

func expandLabelOptions(tfList []interface{}) *types.LabelOptions {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	options := &types.LabelOptions{}

	if v, ok := tfMap["custom_label"].(string); ok {
		options.CustomLabel = aws.String(v)
	}
	if v, ok := tfMap["visibility"].(string); ok {
		options.Visibility = types.Visibility(v)
	}
	if v, ok := tfMap["font_configuration"].([]interface{}); ok && len(v) > 0 {
		options.FontConfiguration = expandFontConfiguration(v)
	}

	return options
}

func expandFontConfiguration(tfList []interface{}) *types.FontConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.FontConfiguration{}

	if v, ok := tfMap["font_color"].(string); ok && v != "" {
		config.FontColor = aws.String(v)
	}
	if v, ok := tfMap["font_decoration"].(string); ok && v != "" {
		config.FontDecoration = types.FontDecoration(v)
	}
	if v, ok := tfMap["font_style"].(string); ok && v != "" {
		config.FontStyle = types.FontStyle(v)
	}
	if v, ok := tfMap["font_size"].([]interface{}); ok && len(v) > 0 {
		config.FontSize = expandFontSize(v)
	}
	if v, ok := tfMap["font_weight"].([]interface{}); ok && len(v) > 0 {
		config.FontWeight = expandFontWeight(v)
	}

	return config
}

func expandFontSize(tfList []interface{}) *types.FontSize {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.FontSize{}

	if v, ok := tfMap["relative"].(string); ok {
		config.Relative = types.RelativeFontSize(v)
	}

	return config
}

func expandFontWeight(tfList []interface{}) *types.FontWeight {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.FontWeight{}

	if v, ok := tfMap["name"].(string); ok {
		config.Name = types.FontWeightName(v)
	}

	return config
}

func expandComparisonFormatConfiguration(tfList []interface{}) *types.ComparisonFormatConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.ComparisonFormatConfiguration{}

	if v, ok := tfMap["number_display_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NumberDisplayFormatConfiguration = expandNumberDisplayFormatConfiguration(v)
	}
	if v, ok := tfMap["percentage_display_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.PercentageDisplayFormatConfiguration = expandPercentageDisplayFormatConfiguration(v)
	}

	return config
}

func expandNumberDisplayFormatConfiguration(tfList []interface{}) *types.NumberDisplayFormatConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.NumberDisplayFormatConfiguration{}

	if v, ok := tfMap["number_scale"].(string); ok && v != "" {
		config.NumberScale = types.NumberScale(v)
	}
	if v, ok := tfMap["prefix"].(string); ok && v != "" {
		config.Prefix = aws.String(v)
	}
	if v, ok := tfMap["suffix"].(string); ok && v != "" {
		config.Suffix = aws.String(v)
	}
	if v, ok := tfMap["decimal_places_configuration"].([]interface{}); ok && len(v) > 0 {
		config.DecimalPlacesConfiguration = expandDecimalPlacesConfiguration(v)
	}
	if v, ok := tfMap["negative_value_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NegativeValueConfiguration = expandNegativeValueConfiguration(v)
	}
	if v, ok := tfMap["null_value_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NullValueFormatConfiguration = expandNullValueFormatConfiguration(v)
	}
	if v, ok := tfMap["separator_configuration"].([]interface{}); ok && len(v) > 0 {
		config.SeparatorConfiguration = expandNumericSeparatorConfiguration(v)
	}

	return config
}

func expandPercentageDisplayFormatConfiguration(tfList []interface{}) *types.PercentageDisplayFormatConfiguration {
	if len(tfList) == 0 || tfList[0] == nil {
		return nil
	}

	tfMap, ok := tfList[0].(map[string]interface{})
	if !ok {
		return nil
	}

	config := &types.PercentageDisplayFormatConfiguration{}

	if v, ok := tfMap["prefix"].(string); ok && v != "" {
		config.Prefix = aws.String(v)
	}
	if v, ok := tfMap["suffix"].(string); ok && v != "" {
		config.Suffix = aws.String(v)
	}
	if v, ok := tfMap["decimal_places_configuration"].([]interface{}); ok && len(v) > 0 {
		config.DecimalPlacesConfiguration = expandDecimalPlacesConfiguration(v)
	}
	if v, ok := tfMap["negative_value_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NegativeValueConfiguration = expandNegativeValueConfiguration(v)
	}
	if v, ok := tfMap["null_value_format_configuration"].([]interface{}); ok && len(v) > 0 {
		config.NullValueFormatConfiguration = expandNullValueFormatConfiguration(v)
	}
	if v, ok := tfMap["separator_configuration"].([]interface{}); ok && len(v) > 0 {
		config.SeparatorConfiguration = expandNumericSeparatorConfiguration(v)
	}

	return config
}

func flattenFormatConfiguration(apiObject *types.FormatConfiguration) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	if apiObject.DateTimeFormatConfiguration != nil {
		tfMap["date_time_format_configuration"] = flattenDateTimeFormatConfiguration(apiObject.DateTimeFormatConfiguration)
	}
	if apiObject.NumberFormatConfiguration != nil {
		tfMap["number_format_configuration"] = flattenNumberFormatConfiguration(apiObject.NumberFormatConfiguration)
	}
	if apiObject.StringFormatConfiguration != nil {
		tfMap["string_format_configuration"] = flattenStringFormatConfiguration(apiObject.StringFormatConfiguration)
	}

	return []interface{}{tfMap}
}

func flattenDateTimeFormatConfiguration(apiObject *types.DateTimeFormatConfiguration) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	if apiObject.DateTimeFormat != nil {
		tfMap["date_time_format"] = aws.ToString(apiObject.DateTimeFormat)
	}
	if apiObject.NullValueFormatConfiguration != nil {
		tfMap["null_value_format_configuration"] = flattenNullValueFormatConfiguration(apiObject.NullValueFormatConfiguration)
	}
	if apiObject.NumericFormatConfiguration != nil {
		tfMap["numeric_format_configuration"] = flattenNumericFormatConfiguration(apiObject.NumericFormatConfiguration)
	}

	return []interface{}{tfMap}
}

func flattenNullValueFormatConfiguration(apiObject *types.NullValueFormatConfiguration) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	if apiObject.NullString != nil {
		tfMap["null_string"] = aws.ToString(apiObject.NullString)
	}

	return []interface{}{tfMap}
}

func flattenNumericFormatConfiguration(apiObject *types.NumericFormatConfiguration) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	if apiObject.CurrencyDisplayFormatConfiguration != nil {
		tfMap["currency_display_format_configuration"] = flattenCurrencyDisplayFormatConfiguration(apiObject.CurrencyDisplayFormatConfiguration)
	}
	if apiObject.NumberDisplayFormatConfiguration != nil {
		tfMap["number_display_format_configuration"] = flattenNumberDisplayFormatConfiguration(apiObject.NumberDisplayFormatConfiguration)
	}
	if apiObject.PercentageDisplayFormatConfiguration != nil {
		tfMap["percentage_display_format_configuration"] = flattenPercentageDisplayFormatConfiguration(apiObject.PercentageDisplayFormatConfiguration)
	}

	return []interface{}{tfMap}
}

func flattenCurrencyDisplayFormatConfiguration(apiObject *types.CurrencyDisplayFormatConfiguration) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	if apiObject.DecimalPlacesConfiguration != nil {
		tfMap["decimal_places_configuration"] = flattenDecimalPlacesConfiguration(apiObject.DecimalPlacesConfiguration)
	}
	if apiObject.NegativeValueConfiguration != nil {
		tfMap["negative_value_configuration"] = flattenNegativeValueConfiguration(apiObject.NegativeValueConfiguration)
	}
	if apiObject.NullValueFormatConfiguration != nil {
		tfMap["null_value_format_configuration"] = flattenNullValueFormatConfiguration(apiObject.NullValueFormatConfiguration)
	}
	tfMap["number_scale"] = types.NumberScale(apiObject.NumberScale)
	if apiObject.Prefix != nil {
		tfMap["prefix"] = aws.ToString(apiObject.Prefix)
	}
	if apiObject.SeparatorConfiguration != nil {
		tfMap["separator_configuration"] = flattenNumericSeparatorConfiguration(apiObject.SeparatorConfiguration)
	}
	if apiObject.Suffix != nil {
		tfMap["suffix"] = aws.ToString(apiObject.Suffix)
	}
	if apiObject.Symbol != nil {
		tfMap["symbol"] = aws.ToString(apiObject.Symbol)
	}

	return []interface{}{tfMap}
}

func flattenDecimalPlacesConfiguration(apiObject *types.DecimalPlacesConfiguration) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	if apiObject.DecimalPlaces != nil {
		tfMap["decimal_places"] = aws.ToInt64(apiObject.DecimalPlaces)
	}

	return []interface{}{tfMap}
}

func flattenNegativeValueConfiguration(apiObject *types.NegativeValueConfiguration) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	tfMap["display_mode"] = types.NegativeValueDisplayMode(apiObject.DisplayMode)

	return []interface{}{tfMap}
}

func flattenNumericSeparatorConfiguration(apiObject *types.NumericSeparatorConfiguration) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	tfMap["decimal_separator"] = types.NumericSeparatorSymbol(apiObject.DecimalSeparator)
	if apiObject.ThousandsSeparator != nil {
		tfMap["thousands_separator"] = flattenThousandSeparatorOptions(apiObject.ThousandsSeparator)
	}

	return []interface{}{tfMap}
}
func flattenThousandSeparatorOptions(apiObject *types.ThousandSeparatorOptions) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	tfMap["symbol"] = types.NumericSeparatorSymbol(apiObject.Symbol)
	tfMap["visibility"] = types.Visibility(apiObject.Visibility)

	return []interface{}{tfMap}
}

func flattenNumberDisplayFormatConfiguration(apiObject *types.NumberDisplayFormatConfiguration) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	if apiObject.DecimalPlacesConfiguration != nil {
		tfMap["decimal_places_configuration"] = flattenDecimalPlacesConfiguration(apiObject.DecimalPlacesConfiguration)
	}
	if apiObject.NegativeValueConfiguration != nil {
		tfMap["negative_value_configuration"] = flattenNegativeValueConfiguration(apiObject.NegativeValueConfiguration)
	}
	if apiObject.NullValueFormatConfiguration != nil {
		tfMap["null_value_format_configuration"] = flattenNullValueFormatConfiguration(apiObject.NullValueFormatConfiguration)
	}
	tfMap["number_scale"] = types.NumberScale(apiObject.NumberScale)
	if apiObject.Prefix != nil {
		tfMap["prefix"] = aws.ToString(apiObject.Prefix)
	}
	if apiObject.SeparatorConfiguration != nil {
		tfMap["separator_configuration"] = flattenNumericSeparatorConfiguration(apiObject.SeparatorConfiguration)
	}
	if apiObject.Suffix != nil {
		tfMap["suffix"] = aws.ToString(apiObject.Suffix)
	}

	return []interface{}{tfMap}
}

func flattenPercentageDisplayFormatConfiguration(apiObject *types.PercentageDisplayFormatConfiguration) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	if apiObject.DecimalPlacesConfiguration != nil {
		tfMap["decimal_places_configuration"] = flattenDecimalPlacesConfiguration(apiObject.DecimalPlacesConfiguration)
	}
	if apiObject.NegativeValueConfiguration != nil {
		tfMap["negative_value_configuration"] = flattenNegativeValueConfiguration(apiObject.NegativeValueConfiguration)
	}
	if apiObject.NullValueFormatConfiguration != nil {
		tfMap["null_value_format_configuration"] = flattenNullValueFormatConfiguration(apiObject.NullValueFormatConfiguration)
	}
	if apiObject.Prefix != nil {
		tfMap["prefix"] = aws.ToString(apiObject.Prefix)
	}
	if apiObject.SeparatorConfiguration != nil {
		tfMap["separator_configuration"] = flattenNumericSeparatorConfiguration(apiObject.SeparatorConfiguration)
	}
	if apiObject.Suffix != nil {
		tfMap["suffix"] = aws.ToString(apiObject.Suffix)
	}

	return []interface{}{tfMap}
}

func flattenNumberFormatConfiguration(apiObject *types.NumberFormatConfiguration) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	if apiObject.FormatConfiguration != nil {
		tfMap["numeric_format_configuration"] = flattenNumericFormatConfiguration(apiObject.FormatConfiguration)
	}

	return []interface{}{tfMap}
}

func flattenStringFormatConfiguration(apiObject *types.StringFormatConfiguration) []interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}
	if apiObject.NullValueFormatConfiguration != nil {
		tfMap["null_value_format_configuration"] = flattenNullValueFormatConfiguration(apiObject.NullValueFormatConfiguration)
	}
	if apiObject.NumericFormatConfiguration != nil {
		tfMap["numeric_format_configuration"] = flattenNumericFormatConfiguration(apiObject.NumericFormatConfiguration)
	}

	return []interface{}{tfMap}
}
