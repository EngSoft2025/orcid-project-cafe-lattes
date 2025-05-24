// services/contact.ts
export type ContactLink = {
  name: string;
  url: string;
  icon: string;
};

export type ContactData = {
  email: string;
  links: ContactLink[];
};

const BASE_URL = 'http://localhost:3001';

export async function getContact(): Promise<ContactData> {
  const res = await fetch(`${BASE_URL}/contact`, { cache: 'no-store' });
  if (!res.ok) throw new Error('Erro ao buscar contato');
  return res.json();
}
