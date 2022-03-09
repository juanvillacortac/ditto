import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Foo } from "../models/foo";
import { FooListFilter } from "../models/foo-list-filter";
import { FooFilter } from "../models/foo-filter";

@Injectable({
  providedIn: "root",
})
export class FooService {
  constructor(
    private _httpClient: HttpClient,
    private _httpHelpersService: HttpHelpersService
  ) {}

  getFooList(filter: FooListFilter) {
    return this._httpClient.get<Foo[]>(
      `/Foo/list`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  getFoo(filter: FooFilter) {
    return this._httpClient.get<Foo>(
      `/Foo/`,
      {
        params: this._httpHelpersService.getHttpParamsFromPlainObject(
          filter,
          false
        ),
      }
    );
  }

  postFoo(model: Foo) {
    return this._httpClient.post<number>(`/Foo/`, model);
  }
}
