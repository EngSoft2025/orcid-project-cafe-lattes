import { RecordDataResponse } from "../types/apiData";


const BASE_API_URL = "http://localhost:8080/api"
export async function fetchRecordData(orcidId: string): Promise<RecordDataResponse> {

    const res = await fetch(`${BASE_API_URL}/searchRecord?orcid_id=${orcidId}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    });

    if (!res.ok) {
        // Lança erro para ser tratado na camada de chamada (ou você pode tratar aqui)
        throw new Error(`Erro ao buscar dados: ${res.status} ${res.statusText}`);
    }

    // Já assume que o JSON retornado corresponde exatamente a RecordDataResponse
    const data = (await res.json()) as RecordDataResponse;
    return data;
}