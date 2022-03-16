import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Foo } from "../models/foo";

@Injectable({
  providedIn: "root",
})
export class FooService {
  constructor(
    private _httpClient: HttpClient,
    private _httpHelpersService: HttpHelpersService
  ) {}
  postFoo(model: Foo) {
    return this._httpClient.post<number>('/foos', model);
  }
}
