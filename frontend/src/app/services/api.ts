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

export interface ResearcherSummary {
    orcid_id: string;
    name: string;
}

export interface ResearcherResults {
researchers: ResearcherSummary[];
num_found: number;
}

export async function fetchResearchersByName(name: string): Promise<ResearcherResults> {
    const res = await fetch(
        `${BASE_API_URL}/searchResearchersByName?name=${encodeURIComponent(name)}`,
            { method: "GET", headers: { "Content-Type": "application/json" } }
        );
        if (!res.ok) throw new Error(`Erro ao buscar por nome: ${res.status}`);
    return res.json() as Promise<ResearcherResults>;
  }