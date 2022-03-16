{{- $listFilter := NodeOption .Model "list-filter" -}}
{{- $filter := NodeOption .Model "filter" -}}
import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { {{ .Model.Name }} } from "../models/{{ .Model.Name | KebabCase }}";
{{- if $listFilter }}
import { {{ $listFilter }} } from "../models/{{ $listFilter | KebabCase }}";
{{- end }}
{{- if $filter }}
import { {{ $filter }} } from "../models/{{ $filter | KebabCase }}";
{{- end }}

@Injectable({
  providedIn: "root",
})
export class {{ .Model.Name }}Service {
  constructor(
    private _httpClient: HttpClient,
    private _httpHelpersService: HttpHelpersService
  ) {}

  {{- if $listFilter }}
  get{{ .Model.Name | Plural }}(filter: {{ $listFilter }}) {
    return this._httpClient.get<{{ .Model.Name }}[]>(
      '/{{ .Model.Name | Plural | KebabCase }}',
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }
  {{ end }}
  {{- if $filter }}
  {{- $filterModel := Model $filter}}
  get{{ .Model.Name }}(filter: {{ $filter }}) {
    {{- $pk := $filterModel.PKProp.Name }}
    const { {{ $pk }} } = filter
    return this._httpClient.get<{{ .Model.Name }}>(`/{{ .Model.Name | Plural | KebabCase }}/${ {{- $pk -}} }`);
  }
  {{ end }}
  post{{ .Model.Name }}(model: {{ .Model.Name }}) {
    return this._httpClient.post<number>('/{{ .Model.Name | Plural | KebabCase }}', model);
  }
}
