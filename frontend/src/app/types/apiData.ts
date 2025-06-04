export interface AffiliationStruct {
  "department-name"?: string | null;
  "role-title": string;
  "start-year"?: string | null;
  "start-month"?: string | null;
  "end-year"?: string | null;
  "end-month"?: string | null;
  "organization-name"?: string | null;
  url?: string | null;
}

export interface PublicationStruct {
  title?: string | null;
  doi?: string | null;
  url?: string | null;
  type: string;
  year?: string | null;
  journal?: string | null;
}

export interface FundingsStruct {
  title?: string | null;
  "start-year"?: string | null;
  "start-month"?: string | null;
  "end-year"?: string | null;
  "end-month"?: string | null;
  "organization-name"?: string | null;
  url?: string | null;
}

export interface RecordDataResponse {
  "orcid-id": string;

  person: {
    "given-names": string;
    "family-names": string;
    "credit-names"?: string | null;
    biography?: string | null;
    contact: {
      email: string[];
      "researcher-url"?: Array<{
        "url-name": string;
        url: string;
      }>;
    };
  };

  keywords: string[];

  affiliations: {
    educations: AffiliationStruct[];
    qualifications: AffiliationStruct[];
    employments: AffiliationStruct[];
    fundings: FundingsStruct[];
  };

  publications: PublicationStruct[];

  MetricsData: {
    PublicationsByYear: Array<{
      year: string; // ex.: "2022"
      count: number;
    }>;
    PublicationsByJournal: Array<{
      title: string;
      count: number;
    }>;
    "h-index": number;
  };
}
