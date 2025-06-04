import { RecordDataResponse } from "@/app/types/apiData";
import { Publicacao } from "@/app/types";


export function extractPublicationsData(record: RecordDataResponse): Publicacao[]{
    const pubs = record.publications ?? [];

    return pubs.map((p, idx)=>{
        //titulo
        const titulo = p.title ?? "";
        
        //Tipo
        const tipo = p.type;

        //ano
        const ano = p.year ? Number(p.year) : 0;

        //Jornal/Conferencia
        const periodico = p.journal ?? ""
        
        //doi
        const doi = p.doi ?? p.url ?? "";

        return{
            id: idx,
            titulo,
            tipo,
            ano,
            periodico,
            doi,
        }
    })
}