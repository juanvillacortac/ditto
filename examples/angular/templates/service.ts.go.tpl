{{- $filter := NodeOption .Model "filter" -}}
import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { {{ .Model.Name }} } from "../models/{{ .Model.Name | KebabCase }}";
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

  {{- if $filter }}
  get{{ .Model.Name | Plural }}(filter: {{ $filter }}) {
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
  {{- $pk := .Model.PKProp }}
  {{- if $pk }}
  get{{ .Model.Name }}({{ $pk.Name }}: {{ $pk.Type }}) {
    return this._httpClient.get<{{ .Model.Name }}>(`/{{ .Model.Name | Plural | KebabCase }}/${ {{- $pk.Name -}} }`);
  }
  {{ end }}
  post{{ .Model.Name }}(model: {{ .Model.Name }}) {
    return this._httpClient.post<number>('/{{ .Model.Name | Plural | KebabCase }}', model);
  }
}
