import { RecordDataResponse } from "@/app/types/apiData";
import { Pesquisador, Afiliacao } from "@/app/types";

export function extractResearcherData(record: RecordDataResponse): Pesquisador{
    
    //orcidid
    const orcidId = record["orcid-id"]

    //nome
    const credit = record.person["credit-names"]
    const given = record.person["given-names"]
    const family = record.person["family-names"]
    const nome = credit?.trim()
        ? credit
        : `${given} ${family}`;
    
    //area principal
    const kws = record.keywords ?? [];
    const areaPrincipal = kws.length > 0 
        ? kws[0]
        : "Sem area principal informada";

    //foto
    const foto = "https://picsum.photos/200" //SEM IMAGEM POR ENQUANTO NA API

    //afiliacoes (usa employments se estiver vazio)
    const rawEmp = record.affiliations.employments ?? [];
    const afiliacoes: Afiliacao[] = rawEmp.map((emp)=>{
        const instituicao =emp["organization-name"] ?? "Instituição não informada";

        const titulo = emp["role-title"] ??  "Cargo não informado";

        // período: formata “MM/YYYY” se houver; se não, “—”
        const startM = emp["start-month"]?.padStart(2, "0");
        const startY = emp["start-year"];
        const endM = emp["end-month"]?.padStart(2, "0");
        const endY = emp["end-year"];

        let periodo = "";
        if (startY) {
        periodo += startM && startM !== "00"
            ? `${startM}/${startY}`
            : `${startY}`;
        }
        if (endY) {
        const endParte = endM && endM !== "00"
            ? `${endM}/${endY}`
            : `${endY}`;
        periodo += ` – ${endParte}`;
        }
        // se nem startY nem endY existirem, deixamos vazio
        if (!periodo) periodo = "Período não disponível";

        const descricao = `${titulo} (${periodo})`;

        return { instituicao, descricao };
    })

    const palavrasChave = kws

    return{
        id: orcidId,
        nome,
        areaPrincipal,
        foto,
        afiliacoes,
        palavrasChave
    }
}