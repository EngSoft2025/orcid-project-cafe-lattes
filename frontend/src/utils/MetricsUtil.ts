
import { RecordDataResponse } from "@/app/types/apiData";
import { MetricsData, YearMetric, VenueMetric } from "@/app/types";

export function extractMetricsData(record: RecordDataResponse): MetricsData {
  const rawYears = record.MetricsData.PublicationsByYear ?? [];
  const rawVenues = record.MetricsData.PublicationsByJournal ?? [];
  const rawHIndex = record.MetricsData["h-index"] ?? 0;


  const publicationsByYear: YearMetric[] = rawYears
    .map(item => ({
      year: Number(item.year),
      count: item.count,
    }))
    .sort((a, b) => a.year - b.year);


  const publicationsByVenue: VenueMetric[] = rawVenues
    .map(item => ({
      name: item.title,
      count: item.count,
    }))
    .sort((a, b) => b.count - a.count);

  const hIndex = rawHIndex;

  return {
    publicationsByYear,
    publicationsByVenue,
    hIndex,
  };
}
