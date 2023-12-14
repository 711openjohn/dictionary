import axios, { AxiosInstance } from "axios";
import { DictionaryPage } from "../types/dictionary";
export class DictionaryClient {
  static apiBaseUrl: string = "https://brain-1-u3445158.deta.app";
  static httpClient: AxiosInstance;
  private static instance = new DictionaryClient();

  constructor() {
    DictionaryClient.httpClient = axios.create({
      headers: { "Content-type": "application/json" },
    });
  }

  public static getInstance(): DictionaryClient {
    return DictionaryClient.instance;
  }

  public async lookup(symbol: string): Promise<DictionaryPage | undefined> {
    return await DictionaryClient.httpClient
      .get(`${DictionaryClient.apiBaseUrl}/dictionary/lookup/?symbol=${symbol}`)
      .then((resp) => {
        return resp.data;
      })
      .catch(() => {
        return undefined;
      });
  }
}
