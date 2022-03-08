import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { AuxiliaryListFilter } from "../models/auxiliary-list-filter";
import { AuxiliaryListFilterListFilter } from "../models/auxiliary-list-filter-list-filter";
import { AuxiliaryListFilterFilter } from "../models/auxiliary-list-filter-filter";

@Injectable({
  providedIn: "root",
})
export class AuxiliaryListFilterService {
  constructor(
    private _httpClient: HttpClient,
    private _httpHelpersService: HttpHelpersService
  ) {}

  getAuxiliaryListFilterList(filter: AuxiliaryListFilterListFilter) {
    return this._httpClient.get<AuxiliaryListFilter[]>(
      `/AuxiliaryListFilter/list`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  getAuxiliaryListFilter(filter: AuxiliaryListFilterFilter) {
    return this._httpClient.get<AuxiliaryListFilter>(
      `/AuxiliaryListFilter/`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  postAuxiliaryListFilter(model: AuxiliaryListFilter) {
    return this._httpClient.post<number>(`/AuxiliaryListFilter/`, model);
  }
}
