# code_helper
code_helper可以协助我们生成代码，用于自动生成一些繁琐且重复性的代码
## How to use
example：
接口文档定义json：
{
"Legal": true,
"Students": [{
"StudentsID": "7124324354",
"Name": "小明",
"Subjects": [{
"Name": "语文",
"Score": 100,
"Percent": 40.00
},
{
"Name": "数学",
"Score": 50,
"Percent": 40.00
}
]
},
{
"StudentsID": "7124324354",
"Name": "小红",
"Subjects": [{
"Name": "语文",
"Score": 100,
"Percent": 40.00
}]
}
]
}

将上述字符串替代main函数中GetGoStructDefine()方法中的参数并执行该程序
命令行输出：
type SubjectsItemModel struct {
Name            string          "json:`Name`"
Score           float64         "json:`Score`"
Percent         float64         "json:`Percent`"
}

type StudentsItemModel struct {
StudentsID              string          "json:`StudentsID`"
Name            string          "json:`Name`"
Subjects                []SubjectsItemModel             "json:`Subjects`"
}

type RootModel struct {
Legal           bool            "json:`Legal`"
Students                []StudentsItemModel             "json:`Students`"
}
