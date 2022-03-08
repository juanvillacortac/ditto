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
