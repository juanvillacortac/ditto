import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { {{ .Model.ModelName }} } from "../models/{{ .Model.ModelName | ToKebabCase }}";

@Injectable({
  providedIn: "root",
})
export class {{ .Model.ModelName }}Service {
  constructor(
    private _httpClient: HttpClient,
    private _httpHelpersService: HttpHelpersService
  ) {}

  get{{ .Model.ModelName }}List() {
    return this._httpClient.get<{{ .Model.ModelName }}[]>(
      `/{{ .Model.ModelName }}/`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filters,
          false
        ),
      }
    );
  }

  get{{ .Model.ModelName }}() {
    return this._httpClient.get<{{ .Model.ModelName }}>(
      `/{{ .Model.ModelName }}/`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filters,
          false
        ),
      }
    );
  }
}
