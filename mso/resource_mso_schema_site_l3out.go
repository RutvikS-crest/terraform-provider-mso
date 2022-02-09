package mso

import (
	"log"
	"strings"

	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceMSOSchemaSiteL3out() *schema.Resource {
	return &schema.Resource{
		Create: resourceMSOSchemaSiteL3outCreate,
		Read:   resourceMSOSchemaSiteL3outRead,
		Delete: resourceMSOSchemaSiteL3outDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMSOSchemaSiteL3outImport,
		},
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"l3out_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"vrf_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"template_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"site_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
			"schema_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 1000),
			},
		},
	}
}

func setMSOSchemaSiteL3outAttributes(l3outMap *models.IntersiteL3outs, d *schema.ResourceData) {
	d.Set("l3out_name", l3outMap.L3outName)
	d.Set("vrf_name", l3outMap.VRFName)
	d.Set("template_name", l3outMap.TemplateName)
	d.Set("site_id", l3outMap.SiteId)
	d.Set("schema_id", l3outMap.SchemaID)
}

func resourceMSOSchemaSiteL3outImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error){
	log.Println("[DEBUG] Schema Site L3out: Beginning Import",d.Id())
	msoClient := m.(*client.Client)
	getAttributes:=strings.Split(d.Id(),"/")
	if 
	log.Println("[DEBUG] Schema Site L3out: Import Completed",d.Id())
}

func resourceMSOSchemaSiteL3outCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Schema Site L3out: Beginning Creation")
	msoClient := m.(*client.Client)
	schemaId := d.Get("schema_id").(string)
	siteId := d.Get("site_id").(string)
	templateName := d.Get("template_name").(string)
	vrfName := d.Get("vrf_name").(string)
	l3outName := d.Get("l3out_name").(string)
	var l3outMap *models.IntersiteL3outs
	l3outMap.L3outName = l3outName
	l3outMap.VRFName = vrfName
	l3outMap.SiteId = siteId
	l3outMap.TemplateName = templateName
	l3outMap.SchemaID = schemaId
	err := msoClient.CreateIntersiteL3outs(l3outMap)
	if err != nil {
		return err
	}
	d.SetId(l3outName)
	log.Printf("[DEBUG] Schema Site L3out: Creation Completed")
	return resourceMSOSchemaSiteL3outRead(d, m)
}

func resourceMSOSchemaSiteL3outRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Schema Site L3out: Beginning Read", d.Id())
	msoClient := m.(*client.Client)
	schemaId := d.Get("schema_id").(string)
	siteId := d.Get("site_id").(string)
	templateName := d.Get("template_name").(string)
	vrfName := d.Get("vrf_name").(string)
	l3outName := d.Get("l3out_name").(string)
	var l3outMap models.IntersiteL3outs
	l3outMap.L3outName = l3outName
	l3outMap.VRFName = vrfName
	l3outMap.SiteId = siteId
	l3outMap.TemplateName = templateName
	l3outMap.SchemaID = schemaId
	l3outMapRemote, err := msoClient.ReadIntersiteL3outs(l3outMap)
	if err != nil {
		d.SetId("")
		return err
	}
	setMSOSchemaSiteL3outAttributes(l3outMapRemote, d)
	d.SetId(l3outName)
	log.Println("[DEBUG] Schema Site L3out: Reading Completed", d.Id())
	return nil
}

func resourceMSOSchemaSiteL3outDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Schema Site L3out: Beginning Destroy", d.Id())
	msoClient := m.(*client.Client)
	schemaId := d.Get("schema_id").(string)
	siteId := d.Get("site_id").(string)
	templateName := d.Get("template_name").(string)
	vrfName := d.Get("vrf_name").(string)
	l3outName := d.Get("l3out_name").(string)
	var l3outMap *models.IntersiteL3outs
	l3outMap.L3outName = l3outName
	l3outMap.VRFName = vrfName
	l3outMap.SiteId = siteId
	l3outMap.TemplateName = templateName
	l3outMap.SchemaID = schemaId
	err := msoClient.DeleteIntersiteL3outs(l3outMap)
	if err != nil {
		return err
	}
	log.Println("[DEBUG] Schema Site L3out: Beginning Destroy", d.Id())
	d.SetId("")
	return err
}
