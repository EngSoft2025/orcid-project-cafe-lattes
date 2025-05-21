//Run JSON-SERVER:  json-server --watch db.json --port 3001

import PublicationTable from "./components/PublicationTable";
import ResearcherResume from "./components/ResearcherResume";

export default function Home() {
  return (
    <div className="bg-background-darker min-h-screen">
      <ResearcherResume />
      <PublicationTable />
    </div>
  );
}
