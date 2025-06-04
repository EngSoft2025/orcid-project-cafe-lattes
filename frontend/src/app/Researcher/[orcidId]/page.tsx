//Run JSON-SERVER:  json-server --watch db.json --port 3001

//componentes
import ContactSection from "../../components/ContactSection";
import Metrics from "../../components/Metrics";
import PublicationTable from "../../components/PublicationTable";
import ResearcherResume from "../../components/ResearcherResume";


//api e utils
import { fetchRecordData } from "@/app/services/api";
import { extractResearcherData } from "@/utils/ResearcherUtil";
import { extractPublicationsData } from "@/utils/PublicationsUtil";
import { extractContactData } from "@/utils/ContactUtil";

//tipos
import { Pesquisador, ContactData, Publicacao } from "@/app/types";


interface ResearcherPageProps{
  params:{
    orcidId: string;
  };
}

export default async function ResearcherPage(props: ResearcherPageProps) {
  const resolvedParams = await Promise.resolve(props.params);

  const {orcidId} = resolvedParams;

  let pesquisador: Pesquisador | null = null;
  let contactData: ContactData | null = null;
  let publicacoes: Publicacao[] = []

  try{
    const record = await fetchRecordData(orcidId)

    //Extraindo dados da requisicao para utilizar nos componentes
    pesquisador = extractResearcherData(record);
    contactData = extractContactData(record)
    publicacoes = extractPublicationsData(record)

  }catch(error){

    contactData = null
  }


  return (
    <div className="bg-background-darker min-h-screen">
      <ResearcherResume pesquisador={pesquisador} />
      <PublicationTable publicacoes={publicacoes} />
      <Metrics />
      <ContactSection contact={contactData}/>
    </div>
  );  
}
