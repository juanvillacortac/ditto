import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Vector } from "../models/vector";

@Injectable({
  providedIn: "root",
})
export class VectorService {
  constructor(
    private _httpClient: HttpClient,
    private _httpHelpersService: HttpHelpersService
  ) {}
  postVector(model: Vector) {
    return this._httpClient.post<number>('/vectors', model);
  }
}
