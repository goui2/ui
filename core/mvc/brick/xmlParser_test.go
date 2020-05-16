package brick

import (
	"fmt"
	"strings"
	"testing"
)

const xml01 = `
 <mvc:View controllerName="testdata.complexsyntax" xmlns:core="sap.ui.core"
 xmlns:mvc="sap.ui.core.mvc" xmlns="sap.ui.commons" xmlns:table="sap.ui.table"
 xmlns:html="http:www.w3.org/1999/xhtml">
 <html:h2>
  	<Label text="Hello Mr. {path:'/singleEntry/firstName', formatter:'.myFormatter'}, {/singleEntry/lastName}"></Label>
 </html:h2>
 <table:Table rows="{/table}">
	<table:columns>
		<table:Column>
			<Label text="Name"></Label>
			<table:template>
				<TextField value="{path:'gender', formatter:'.myGenderFormatter'} {firstName}, {lastName}"></TextField>
			</table:template>
		</table:Column>
		<table:Column>
			<Label text="Birthday"></Label>
			<table:template>
				<TextField value="{parts:[{path:'birthday/day'},{path:'birthday/month'},{path:'birthday/year'}], formatter:'my.globalFormatter'}"></TextField>
			</table:template>
		</table:Column>
	 </table:columns>
  </table:Table>
  <html:h2>
  	<Label text="A type test: {path:'/singleEntry/amount', type:'sap.ui.model.type.Float', formatOptions: { minFractionDigits: 1}} EUR"></Label>
  </html:h2>
 </mvc:View>
`

func TestXmlParser(t *testing.T) {
	r := strings.NewReader(xml01)
	v := ParseXmlView(r)
	fmt.Printf("%#v", v)
}
