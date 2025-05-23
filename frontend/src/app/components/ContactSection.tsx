'use client';

import { useEffect, useState } from 'react';
import { getContact, ContactData } from '@/app/services/contact';
import { ArrowRight } from 'lucide-react';

export default function ContactSection() {
  const [contact, setContact] = useState<ContactData | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    async function fetchContact() {
      try {
        const data = await getContact();
        setContact(data);
      } catch (error) {
        console.error('Erro ao buscar contato:', error);
        setContact(null);
      } finally {
        setIsLoading(false);
      }
    }

    fetchContact();
  }, []);

  const getIconForLink = (name: string) => {
    name = name.toLowerCase();

    if (name.includes('orcid')) return '/images/orcidImage.png';
    if (name.includes('scholar')) return '/images/scholar.png';
    if (name.includes('researchgate')) return '/images/researchgate.png';
    if (name.includes('linkedin')) return '/images/linkedin.png';
    if (name.includes('github')) return '/images/github.png';

    return '/images/baseImage.png'; // Ícone padrão
  };

  if (isLoading) {
    return (
      <section id="contact" className="bg-background-darker px-4 sm:px-[5%] lg:px-[15%] py-10">
        <h2 className="text-xl font-semibold relative before:absolute before:left-0 before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-6 before:bg-primary-yellow pl-4">
          Contato e Links
        </h2>
        <div className="max-w-md mx-auto">
          <div className="h-12 mb-4 bg-background rounded animate-pulse" />
          <div className="space-y-3">
            {Array(3)
              .fill(0)
              .map((_, i) => (
                <div key={i} className="h-14 bg-background rounded animate-pulse" />
              ))}
          </div>
        </div>
      </section>
    );
  }

  if (!contact) {
    return (
      <section id="contact" className="bg-background-darker px-4 sm:px-[5%] lg:px-[15%] py-10">
        <h2 className="text-xl font-semibold relative before:absolute before:left-0 before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-6 before:bg-primary-yellow pl-4">
          Contato e Links
        </h2>
        <div className="text-center py-12">
          <p className="text-gray-400">Nenhuma informação de contato disponível.</p>
        </div>
      </section>
    );
  }

  return (
    <section id="contact" className="bg-background-darker px-4 sm:px-[5%] lg:px-[15%] py-10">
      <h2 className="text-xl font-semibold relative before:absolute before:left-0 before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-6 before:bg-primary-yellow pl-4">
        Contato e Links
      </h2>

      <div className="max-w-md mx-auto mt-20">
        {contact.email && (
          <div className="bg-background border-border-color flex items-center mb-4 rounded-lg border bg-card text-card-foreground shadow-sm">
            <div className="py-6 px-5">
              <div className="flex items-center justify-between">
                <span className="text-gray-400 pr-3">Email:</span>
                <a
                  href={`mailto:${contact.email}`}
                  className="text-primary-yellow hover:underline"
                >
                  {contact.email}
                </a>
              </div>
            </div>
          </div>
        )}

        {contact.links.length > 0 && (
          <div className="space-y-3">
            {contact.links.map((link, index) => (
              <div
                key={index}
                className="bg-background border-border-color flex items-center mb-4 rounded-lg border bg-card text-card-foreground shadow-sm"
              >
                <div className="py-6 px-5 flex">
                  <a
                    href={link.url}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex items-center gap-2 justify-between hover:text-primary-yellow transition-colors"
                  >
                    <div className="flex items-center">
                      <img
                        src={getIconForLink(link.name)}                        alt={link.name}
                        className="h-5 w-5 mr-3"
                      />
                      <span>{link.name}</span>
                    </div>
                    <span className="text-primary-yellow">
                      <ArrowRight size={16}/>
                    </span>
                  </a>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </section>
  );
}