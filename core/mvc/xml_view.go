package mvc

import (
	"io"

	"github.com/goui2/ui/base"
	"github.com/goui2/ui/core/mvc/brick"
)

const (
	pfxMvc = "goui.core.mvc"
)

// <mvc:View controllerName="testdata.complexsyntax" xmlns:core="sap.ui.core"
// xmlns:mvc="sap.ui.core.mvc" xmlns="sap.ui.commons" xmlns:table="sap.ui.table"
// xmlns:html="http://www.w3.org/1999/xhtml">
// <html:h2>
//  	<Label text="Hello Mr. {path:'/singleEntry/firstName', formatter:'.myFormatter'}, {/singleEntry/lastName}"></Label>
// </html:h2>
// <table:Table rows="{/table}">
//	<table:columns>
//		<table:Column>
//			<Label text="Name"></Label>
//			<table:template>
//				<TextField value="{path:'gender', formatter:'.myGenderFormatter'} {firstName}, {lastName}"></TextField>
//			</table:template>
//		</table:Column>
//		<table:Column>
//			<Label text="Birthday"></Label>
//			<table:template>
//				<TextField value="{parts:[{path:'birthday/day'},{path:'birthday/month'},{path:'birthday/year'}], formatter:'my.globalFormatter'}"></TextField>
//			</table:template>
//		</table:Column>
//	 </table:columns>
//  </table:Table>
//  <html:h2>
//  	<Label text="A type test: {path:'/singleEntry/amount', type:'sap.ui.model.type.Float', formatOptions: { minFractionDigits: 1}} EUR"></Label>
//  </html:h2>
// </mvc:View>

type XMLView interface {
	View
}

type xmlView struct {
	View
	//	controller Controller
}

func constructorXmlView(md base.MetaData, parent base.Constructor) base.Constructor {
	return func(id string, s ...base.InstanceSetting) base.Object {
		mo := &xmlView{}
		parentSettings := s //append(s, core.ISRenderer{mo.renderer})
		parentSettings = base.AdjustSelfSetting(mo, md, parentSettings)
		mo.View = parent.New(id, parentSettings...).(View)

		return mo
	}
}

func NewXMLView(name string, r io.Reader) ViewCreator {
	viewBrick := brick.ParseXmlView(r)
	if viewBrick.ClassName() != MD_View.GetName() {
		panic("unknown view type " + viewBrick.ClassName())
	}

	viewCtrl := traverseViewBrick(viewBrick, MD_View)

	return viewCtrl
}
