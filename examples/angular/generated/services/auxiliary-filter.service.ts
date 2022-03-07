import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { AuxiliaryFilter } from "../models/auxiliary-filter";
import { AuxiliaryFilterListFilter } from "../models/auxiliary-filter-list-filter";
import { AuxiliaryFilterFilter } from "../models/auxiliary-filter-filter";

@Injectable({
  providedIn: "root",
})
export class AuxiliaryFilterService {
  constructor(
    private _httpClient: HttpClient,
    private _httpHelpersService: HttpHelpersService
  ) {}

  getAuxiliaryFilterList(filter: AuxiliaryFilterListFilter) {
    return this._httpClient.get<AuxiliaryFilter[]>(
      `/AuxiliaryFilter/list`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  getAuxiliaryFilter(filter: AuxiliaryFilterFilter) {
    return this._httpClient.get<AuxiliaryFilter>(
      `/AuxiliaryFilter/`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  postAuxiliaryFilter(model: AuxiliaryFilter) {
    return this._httpClient.post<number>(`/AuxiliaryFilter/`, model);
  }
}
