export default function HomePage() {
    return (
        <main className="w-full min-h-[60vh] bg-background-darker text-foreground">
        <section className="px-4 sm:px-[5%] lg:px-[10%] py-6 w-full flex items-center flex-col gap-10">
            <h1 className="text-5xl font-bold text-center w-[70%]">
                <span className="text-primary-yellow">C</span>
                onsulta e 
                <span className="text-primary-yellow"> L</span>
                eitura
                <span className="text-primary-yellow"> A</span>
                vançada de 
                <span className="text-primary-yellow"> R</span>
                esultados 
                <span className="text-primary-yellow"> A</span>
                cadêmicos
            </h1>

            <p className="text-center text-lg mt-4 text-clara-gray">
                Busque um pesquisador pelo ORCID iD para visualizar seu perfil, publicações, colaborações e métricas acadêmicas.
            </p>

            <div className="bg-background w-[40%] rounded-md border-border-color border p-6">
                <h2 className="text-xl text-center font-semibold">
                    Como usar:
                </h2>
                <ol className="list-decimal list-inside text-clara-gray mt-2 marker:text-primary-yellow marker:font-semibold flex flex-col gap-5">
                    <li>Digite um ORCID iD no formato xxxx-xxxx-xxxx-xxxx no campo de busca.</li>
                    <li>Clique em "Buscar" para visualizar as informações do pesquisador.</li>
                    <li>Navegue pelas diferentes seções para explorar publicações, colaborações e métricas.</li>
                </ol>
                <p className="text-gray-600 text-xs text-center mt-5">Experimente com o ORCID ID: 0000-0002-1234-5678</p>
            </div>

        </section>
        </main>
    );
}