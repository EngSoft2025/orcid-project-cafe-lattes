type SingleDataCardProps = {
  title: string;
  value: number;
};

export default function SingleDataCard({ title, value}: SingleDataCardProps) {
  return (
      <div className="bg-background rounded-lg border border-border-color p-6">
        <h3 className="text-gray-400 mb-2">{title}</h3>
        <p className="text-white text-3xl font-bold">{value}</p>
      </div>
  );
}