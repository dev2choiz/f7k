package annotation

type IAnnotation interface {
	Name() 				 							string
	SetName(name string) 							IAnnotation
	DataAssoc() 									map[string]interface{}
	SetDataAssoc(dataAssoc map[string]interface{}) 	IAnnotation
	Data() 											[]interface{}
	SetData(data []interface{}) 					IAnnotation
	PlainText() 									string
	SetPlainText(plainText string) 					IAnnotation
}

type Annotation struct {
	name      string
	dataAssoc map[string]interface{}
	data      []interface{}
	plainText string
}

func (a *Annotation) Name() string {
	return a.name
}

func (a *Annotation) SetName(name string) IAnnotation {
	a.name = name

	return a
}

func (a *Annotation) DataAssoc() map[string]interface{} {
	return a.dataAssoc
}

func (a *Annotation) SetDataAssoc(dataAssoc map[string]interface{}) IAnnotation {
	a.dataAssoc = dataAssoc

	return a
}

func (a *Annotation) Data() []interface{} {
	return a.data
}

func (a *Annotation) SetData(data []interface{}) IAnnotation {
	a.data = data

	return a
}

func (a *Annotation) PlainText() string {
	return a.plainText
}

func (a *Annotation) SetPlainText(plainText string) IAnnotation {
	a.plainText = plainText

	return a
}
