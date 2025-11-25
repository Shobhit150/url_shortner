
import React from 'react'
import AnalyticsViewer from "@/components/AnalyticsViewer";

export default function Page({ params }: { params: { slug: string } }) {
  return <AnalyticsViewer initialSlug={params.slug} />;
}