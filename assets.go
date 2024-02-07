package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsf956610d26c556e089efaed838d63cd1078c1c03 = "<!doctype html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <title>Seaii FileServer</title>\n    <link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css\" rel=\"stylesheet\" integrity=\"sha384-GLhlTQ8iRABdZLl6O3oVMWSktQOp6b7In1Zl3/Jr59b6EGGoI1aFkw7cmDA6j6gD\" crossorigin=\"anonymous\">\n</head>\n<body>\n\n<nav class=\"navbar bg-dark navbar-expand-lg\" data-bs-theme=\"dark\">\n    <div class=\"container-fluid\">\n        <a class=\"navbar-brand\" href=\"/file\">Navbar</a>\n        <button class=\"navbar-toggler\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#navbarNavAltMarkup\" aria-controls=\"navbarNavAltMarkup\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\n            <span class=\"navbar-toggler-icon\"></span>\n        </button>\n        <div class=\"collapse navbar-collapse\" id=\"navbarNavAltMarkup\">\n            <div class=\"navbar-nav\">\n                <a class=\"nav-link active\" aria-current=\"page\" href=\"/file\">Home</a>\n                <a class=\"nav-link\" href=\"/upload\">Upload</a>\n                <a class=\"nav-link\" href=\"/logs\">Logs</a>\n            </div>\n        </div>\n    </div>\n</nav>\n\n<form action=\"/upload\" method=\"post\" enctype=\"multipart/form-data\">\n    <div class=\"mb-3\">\n        <label for=\"formFileMultiple\" class=\"form-label\">上传多个文件</label>\n        <input class=\"form-control\" type=\"file\" name=\"files\" id=\"formFileMultiple\" multiple>\n    </div>\n    <div class=\"col-auto\">\n        <button type=\"submit\" class=\"btn btn-primary mb-3\">上 传</button>\n    </div>\n</form>\n\n<script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-/mhDoLbDldZc3qpsJHpLogda//BVZbgYuw6kof4u2FrCedxOtgRZDTHgHUhOCVim\" crossorigin=\"anonymous\"></script>\n</body>\n</html>"
var _Assetsac6fc5057113e44dc39b273d2d43195b7a483751 = "<!doctype html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <title>Seaii FileServer</title>\n    <link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css\" rel=\"stylesheet\" integrity=\"sha384-GLhlTQ8iRABdZLl6O3oVMWSktQOp6b7In1Zl3/Jr59b6EGGoI1aFkw7cmDA6j6gD\" crossorigin=\"anonymous\">\n</head>\n<body>\n\n<nav class=\"navbar bg-dark navbar-expand-lg\" data-bs-theme=\"dark\">\n    <div class=\"container-fluid\">\n        <a class=\"navbar-brand\" href=\"/file\">Navbar</a>\n        <button class=\"navbar-toggler\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#navbarNavAltMarkup\" aria-controls=\"navbarNavAltMarkup\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\n            <span class=\"navbar-toggler-icon\"></span>\n        </button>\n        <div class=\"collapse navbar-collapse\" id=\"navbarNavAltMarkup\">\n            <div class=\"navbar-nav\">\n                <a class=\"nav-link active\" aria-current=\"page\" href=\"/file\">Home</a>\n                <a class=\"nav-link\" href=\"/upload\">Upload</a>\n                <a class=\"nav-link\" href=\"/logs\">Logs</a>\n            </div>\n        </div>\n    </div>\n</nav>\n\n<div class=\"list-group\">\n    {{range $file := .files}}\n    <a href=\"/download/{{$file.Name}}\" class=\"list-group-item list-group-item-action list-group-item-secondary\">\n        <div class=\"d-flex w-100 justify-content-between\">\n            <h6 class=\"mb-1\">{{$file.Name}}</h6>\n            <small class=\"text-muted\">{{$file.Time}}</small>\n        </div>\n    </a>\n    {{end}}\n</div>\n\n<script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-/mhDoLbDldZc3qpsJHpLogda//BVZbgYuw6kof4u2FrCedxOtgRZDTHgHUhOCVim\" crossorigin=\"anonymous\"></script>\n</body>\n</html>"
var _Assetsf06e358943557ea1ff48cf37ea660d128f913380 = "<!doctype html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <title>Seaii FileServer</title>\n    <link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css\" rel=\"stylesheet\" integrity=\"sha384-GLhlTQ8iRABdZLl6O3oVMWSktQOp6b7In1Zl3/Jr59b6EGGoI1aFkw7cmDA6j6gD\" crossorigin=\"anonymous\">\n</head>\n<body>\n\n<nav class=\"navbar bg-dark navbar-expand-lg\" data-bs-theme=\"dark\">\n    <div class=\"container-fluid\">\n        <a class=\"navbar-brand\" href=\"/file\">Navbar</a>\n        <button class=\"navbar-toggler\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#navbarNavAltMarkup\" aria-controls=\"navbarNavAltMarkup\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\n            <span class=\"navbar-toggler-icon\"></span>\n        </button>\n        <div class=\"collapse navbar-collapse\" id=\"navbarNavAltMarkup\">\n            <div class=\"navbar-nav\">\n                <a class=\"nav-link active\" aria-current=\"page\" href=\"/file\">Home</a>\n                <a class=\"nav-link\" href=\"/upload\">Upload</a>\n                <a class=\"nav-link\" href=\"/logs\">Logs</a>\n            </div>\n        </div>\n    </div>\n</nav>\n\n<ul class=\"list-group list-group-flush\">\n    {{range $log := .logs}}\n    <li class=\"list-group-item\">\n        <small>{{$log}}</small>\n    </li>\n    {{end}}\n</ul>\n\n\n<script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-/mhDoLbDldZc3qpsJHpLogda//BVZbgYuw6kof4u2FrCedxOtgRZDTHgHUhOCVim\" crossorigin=\"anonymous\"></script>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"file.tmpl", "logs.tmpl", "upload.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1707277791, 1707277791926235900),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1707276494, 1707276494224491200),
		Data:     nil,
	}, "/templates/file.tmpl": &assets.File{
		Path:     "/templates/file.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1707276494, 1707276494224491200),
		Data:     []byte(_Assetsac6fc5057113e44dc39b273d2d43195b7a483751),
	}, "/templates/logs.tmpl": &assets.File{
		Path:     "/templates/logs.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1707276494, 1707276494224491200),
		Data:     []byte(_Assetsf06e358943557ea1ff48cf37ea660d128f913380),
	}, "/templates/upload.tmpl": &assets.File{
		Path:     "/templates/upload.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1707276494, 1707276494224491200),
		Data:     []byte(_Assetsf956610d26c556e089efaed838d63cd1078c1c03),
	}}, "")
