# Rosetta ✨
An extensible CLI tool for generic code generation based in Protobuff and YAML source files

## Why?
Because in my actual work we have a tedious developing workflow. We don't use ORMs and we adopt a "microservices" project structure try.
I develop this tool for generate C# microservices, Angular models and services and SQLServer SPs based in a user given schema in a single shot.

## How this works?

For generate code we need a ".json" source config with a estructure like:

```jsonc
{
  "schema": "schema.proto", // or "schema.yml"
  "generators": [
    {
      "name": "TS Models",
      "template": "templates/model.ts.go.tpl",
      "output": "generated/models/[model-].ts",
      "types": {
        "int32": "number",
        "int64": "number",
        "double": "number",
        "boolean": "boolean",
        "string": "string",
        "date": "Date"
      },
      "helpers": {
        "\"now()\"": "new Date()"
      }
    },
    {
      "name": "TS Service",
      "template": "templates/service.ts.go.tpl",
      "output": "generated/services/[model-].service.ts",
      "types": {
        "int32": "number",
        "int64": "number",
        "double": "number",
        "boolean": "boolean",
        "string": "string",
        "date": "Date"
      },
      "helpers": {
        "\"now()\"": "new Date()"
      }
    }
  ]
}
```

The json have in first level two fields: `file` and `generators`. In file we define the protobuffer path, and in `generators` we place our generators.

A generator have a `name`, a `template` file path, an `output` file path (with custom helpers for place the current model name on filename), an output `types` map, where the key is the protobuff defined type and value is the output type target, and finally an `helpers` map, where the key is a sentence than be replaced in models default values for the corresponding value in the map.

The proto file is like:
```proto
package ejemplo;

import "google/protobuf/descriptor.proto";

extend google.protobuf.MessageOptions {
  string table_name = 5001;
  string sp_name = 5002;
}

message Auxiliary {
  int64 id = 1;
  string name = 2;
  boolean active = 3;
  int64 idUser = 4;
  string userName = 5;
}

message AuxiliaryListFilter {
  int64 id = 1 [default = -1];
  int32 active = 3 [default = -1];
  int64 idUser = 4 [default = -1];
}

message AuxiliaryFilter {
  int64 id = 1;
}
```

The YAML equivalent:
```yaml
Auxiliary:
  id: int64
  name: string
  active: boolean
  idUser: int64
  userName: string

AuxiliaryListFilter:
  id: { type: int64, default: -1 }
  active: { type: int32, default: -1 }
  idUser: { type: int64, default: -1 }

AuxiliaryFilter:
  id: int64

Vector:
  x: int64
  y: int64

Foo:
  bar: Auxiliary
  vectors:
    type: Vector
    array: true
```

And the templates are:
```
// models.ts.go.tpl

{{- range $key, $dep := .Deps -}}
import { {{ $dep }} } from './{{ $dep | ToKebabCase }}'
{{ end }}
export class {{ .Model.ModelName }} {
  {{- range .Model.Props}}
  {{ .Name }}: {{ .Type }}{{ if .IsArray }}[]{{ end }}{{ if HaveDefaultValue . }} = {{ PropDefaultValue . }}{{ end }}
  {{- end }}
}
```

```
// service.ts.go.tpl

import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { {{ .Model.ModelName }} } from "../models/{{ .Model.ModelName | ToKebabCase }}";
import { {{ .Model.ModelName }}ListFilter } from "../models/{{ .Model.ModelName | ToKebabCase }}-list-filter";
import { {{ .Model.ModelName }}Filter } from "../models/{{ .Model.ModelName | ToKebabCase }}-filter";

@Injectable({
  providedIn: "root",
})
export class {{ .Model.ModelName }}Service {
  constructor(
    private _httpClient: HttpClient,
    private _httpHelpersService: HttpHelpersService
  ) {}

  get{{ .Model.ModelName }}List(filter: {{ .Model.ModelName }}ListFilter) {
    return this._httpClient.get<{{ .Model.ModelName }}[]>(
      `/{{ .Model.ModelName }}/list`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  get{{ .Model.ModelName }}(filter: {{ .Model.ModelName }}Filter) {
    return this._httpClient.get<{{ .Model.ModelName }}>(
      `/{{ .Model.ModelName }}/`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  post{{ .Model.ModelName }}(model: {{ .Model.ModelName }}) {
    return this._httpClient.post<number>(`/{{ .Model.ModelName }}/`, model);
  }
}
```

Finally, we obtain a tree of files:
```
- generated/models/auxiliary-filter.ts
- generated/models/auxiliary-list-filter.ts
- generated/models/auxiliary.ts

- generated/services/auxiliary-filter.service.ts (dummy)
- generated/services/auxiliary-list-filter.service.ts (dummy)
- generated/services/auxiliary.service.ts
```

Content of the `auxiliary.ts` file:

```typescript

export class Auxiliary {
  id: number
  name: string
  active: boolean
  idUser: number
  userName: string
}
```

Content of the `auxiliary.service.ts` file: 

```typescript
import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Auxiliary } from "../models/auxiliary";
import { AuxiliaryListFilter } from "../models/auxiliary-list-filter";
import { AuxiliaryFilter } from "../models/auxiliary-filter";

@Injectable({
  providedIn: "root",
})
export class AuxiliaryService {
  constructor(
    private _httpClient: HttpClient,
    private _httpHelpersService: HttpHelpersService
  ) {}

  getAuxiliaryList(filter: AuxiliaryListFilter) {
    return this._httpClient.get<Auxiliary[]>(
      `/Auxiliary/list`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  getAuxiliary(filter: AuxiliaryFilter) {
    return this._httpClient.get<Auxiliary>(
      `/Auxiliary/`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  postAuxiliary(model: Auxiliary) {
    return this._httpClient.post<number>(`/Auxiliary/`, model);
  }
}
```
