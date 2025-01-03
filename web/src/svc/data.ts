import { CommonResponse, Data } from "../models/types";
import { API_URL } from "../utils/constants";
import { getMethodDeleteHeader, getMethodGetHeader, getMethodPostHeader, getMethodPutHeader } from "../utils/https";

const baseUrl = API_URL + '/data';

export async function getAllData(): Promise<Data[]> {
    const headers = getMethodGetHeader();
    const response = await fetch(baseUrl, headers);
    const data: Data[] = await response.json()
    return data;
}

export async function addData(data: Data): Promise<CommonResponse> {
    const headers = getMethodPostHeader(data)
    const response = await fetch(baseUrl, headers);
    const json: CommonResponse = await response.json()
    return json;
}

export async function updateData(data: Data): Promise<CommonResponse> {
    const headers = getMethodPutHeader(data)
    const response = await fetch(baseUrl, headers);
    const json: CommonResponse = await response.json()
    return json;
}

export async function deleteData(id: string): Promise<CommonResponse> {
    const url = baseUrl + `/${id}`;
    const headers = getMethodDeleteHeader()
    const response = await fetch(url, headers);
    const data: CommonResponse = await response.json()
    return data;
}