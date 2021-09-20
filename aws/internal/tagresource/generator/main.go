// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"

	tftags "github.com/hashicorp/terraform-provider-aws/aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

var (
	serviceName = flag.String("servicename", "", "lowercase service name")
)

type TemplateData struct {
	ServiceName string
}

func main() {
	flag.Parse()

	if len(*serviceName) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	resourceName := fmt.Sprintf("aws_%s_tag", *serviceName)
	resourceFilename := fmt.Sprintf("resource_%s_gen.go", resourceName)
	resourceTestFilename := fmt.Sprintf("resource_%s_gen_test.go", resourceName)

	templateData := TemplateData{
		ServiceName: *serviceName,
	}
	templateFuncMap := template.FuncMap{
		"IdentifierAttributeName": tftags.ServiceIdentifierAttributeName,
		"Title":                   strings.Title,
	}

	if err := generateTemplateFile(resourceFilename, resourceTemplateBody, templateFuncMap, templateData); err != nil {
		log.Fatal(err)
	}

	if err := generateTemplateFile(resourceTestFilename, resourceTestTemplateBody, templateFuncMap, templateData); err != nil {
		log.Fatal(err)
	}
}

func generateTemplateFile(filename string, templateBody string, templateFuncs template.FuncMap, templateData interface{}) error {
	tmpl, err := template.New(filename).Funcs(templateFuncs).Parse(templateBody)

	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templateData)

	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	generatedFileContents, err := format.Source(buffer.Bytes())

	if err != nil {
		return fmt.Errorf("error formatting generated file: %w", err)
	}

	f, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("error creating file (%s): %w", filename, err)
	}

	defer f.Close()

	_, err = f.Write(generatedFileContents)

	if err != nil {
		return fmt.Errorf("error writing to file (%s): %w", filename, err)
	}

	return nil
}

const (
	resourceTemplateBody = `
// Code generated by internal/tagresource/generator/main.go; DO NOT EDIT.

package aws

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/{{ .ServiceName }}"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/aws/internal/tags"
	tftags "github.com/hashicorp/terraform-provider-aws/aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/aws/internal/tfresource"
)

func resourceAws{{ .ServiceName | Title }}Tag() *schema.Resource {
	return &schema.Resource{
		Create: resourceAws{{ .ServiceName | Title }}TagCreate,
		Read:   resourceAws{{ .ServiceName | Title }}TagRead,
		Update: resourceAws{{ .ServiceName | Title }}TagUpdate,
		Delete: resourceAws{{ .ServiceName | Title }}TagDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"{{ .ServiceName | IdentifierAttributeName }}": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAws{{ .ServiceName | Title }}TagCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).{{ .ServiceName }}conn

	identifier := d.Get("{{ .ServiceName | IdentifierAttributeName }}").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)

	{{ if eq .ServiceName "ec2" }}
	if err := tftags.{{ .ServiceName | Title }}CreateTags(conn, identifier, map[string]string{key: value}); err != nil {
	{{- else }}
	if err := tftags.{{ .ServiceName | Title }}UpdateTags(conn, identifier, nil, map[string]string{key: value}); err != nil {
	{{- end }}
		return fmt.Errorf("error creating %s resource (%s) tag (%s): %w", {{ .ServiceName }}.ServiceID, identifier, key, err)
	}

	d.SetId(tftags.SetResourceID(identifier, key))

	return resourceAws{{ .ServiceName | Title }}TagRead(d, meta)
}

func resourceAws{{ .ServiceName | Title }}TagRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).{{ .ServiceName }}conn
	identifier, key, err := tftags.GetResourceID(d.Id())

	if err != nil {
		return err
	}

	value, err := tftags.{{ .ServiceName | Title }}GetTag(conn, identifier, key)

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] %s resource (%s) tag (%s) not found, removing from state", {{ .ServiceName }}.ServiceID, identifier, key)
		d.SetId("")
		return nil
	}

	if err != nil {
		return fmt.Errorf("error reading %s resource (%s) tag (%s): %w", {{ .ServiceName }}.ServiceID, identifier, key, err)
	}

	d.Set("{{ .ServiceName | IdentifierAttributeName }}", identifier)
	d.Set("key", key)
	d.Set("value", value)

	return nil
}

func resourceAws{{ .ServiceName | Title }}TagUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).{{ .ServiceName }}conn
	identifier, key, err := tftags.GetResourceID(d.Id())

	if err != nil {
		return err
	}

	if err := tftags.{{ .ServiceName | Title }}UpdateTags(conn, identifier, nil, map[string]string{key: d.Get("value").(string)}); err != nil {
		return fmt.Errorf("error updating %s resource (%s) tag (%s): %w", {{ .ServiceName }}.ServiceID, identifier, key, err)
	}

	return resourceAws{{ .ServiceName | Title }}TagRead(d, meta)
}

func resourceAws{{ .ServiceName | Title }}TagDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).{{ .ServiceName }}conn
	identifier, key, err := tftags.GetResourceID(d.Id())

	if err != nil {
		return err
	}

	if err := tftags.{{ .ServiceName | Title }}UpdateTags(conn, identifier, map[string]string{key: d.Get("value").(string)}, nil); err != nil {
		return fmt.Errorf("error deleting %s resource (%s) tag (%s): %w", {{ .ServiceName }}.ServiceID, identifier, key, err)
	}

	return nil
}
`
	resourceTestTemplateBody = `
// Code generated by internal/tagresource/generator/main.go; DO NOT EDIT.

package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/{{ .ServiceName }}"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tftags "github.com/hashicorp/terraform-provider-aws/aws/internal/tags"
	tftags "github.com/hashicorp/terraform-provider-aws/aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/aws/internal/tfresource"
)

func testAccCheck{{ .ServiceName | Title }}TagDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).{{ .ServiceName }}conn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_{{ .ServiceName }}_tag" {
			continue
		}

		identifier, key, err := tftags.GetResourceID(rs.Primary.ID)

		if err != nil {
			return err
		}

		_, err = tftags.{{ .ServiceName | Title }}GetTag(conn, identifier, key)

		if tfresource.NotFound(err) {
			continue
		}

		if err != nil {
			return err
		}

		return fmt.Errorf("%s resource (%s) tag (%s) still exists", {{ .ServiceName }}.ServiceID, identifier, key)
	}

	return nil
}

func testAccCheck{{ .ServiceName | Title }}TagExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("%s: missing resource ID", resourceName)
		}

		identifier, key, err := tftags.GetResourceID(rs.Primary.ID)

		if err != nil {
			return err
		}

		conn := testAccProvider.Meta().(*AWSClient).{{ .ServiceName }}conn

		_, err = tftags.{{ .ServiceName | Title }}GetTag(conn, identifier, key)

		if err != nil {
			return err
		}

		return nil
	}
}
`
)
