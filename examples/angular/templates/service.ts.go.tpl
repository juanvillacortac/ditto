import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { {{ .Model.Name }} } from "../models/{{ .Model.Name | KebabCase }}";
import { {{ .Model.Name }}ListFilter } from "../models/{{ .Model.Name | KebabCase }}-list-filter";
import { {{ .Model.Name }}Filter } from "../models/{{ .Model.Name | KebabCase }}-filter";

@Injectable({
  providedIn: "root",
})
export class {{ .Model.Name }}Service {
  constructor(
    private _httpClient: HttpClient,
    private _httpHelpersService: HttpHelpersService
  ) {}

  get{{ .Model.Name }}List(filter: {{ .Model.Name }}ListFilter) {
    return this._httpClient.get<{{ .Model.Name }}[]>(
      `/{{ .Model.Name }}/list`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  get{{ .Model.Name }}(filter: {{ .Model.Name }}Filter) {
    return this._httpClient.get<{{ .Model.Name }}>(
      `/{{ .Model.Name }}/`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  post{{ .Model.Name }}(model: {{ .Model.Name }}) {
    return this._httpClient.post<number>(`/{{ .Model.Name }}/`, model);
  }
}
