import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Vector } from "../models/vector";
import { VectorListFilter } from "../models/vector-list-filter";
import { VectorFilter } from "../models/vector-filter";

@Injectable({
  providedIn: "root",
})
export class VectorService {
  constructor(
    private _httpClient: HttpClient,
    private _httpHelpersService: HttpHelpersService
  ) {}

  getVectorList(filter: VectorListFilter) {
    return this._httpClient.get<Vector[]>(
      `/Vector/list`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  getVector(filter: VectorFilter) {
    return this._httpClient.get<Vector>(
      `/Vector/`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  postVector(model: Vector) {
    return this._httpClient.post<number>(`/Vector/`, model);
  }
}
