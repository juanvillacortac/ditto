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
