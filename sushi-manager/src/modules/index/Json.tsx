import Header from "../../components/typography/Header";
import JsonView from "react18-json-view";

interface JsonProps {
  data: any;
}

function Json({ data }: JsonProps) {
  return (
    <div>
      <Header text="Json" align="left" size="md" />
      <div className="bg-neutral-200 px-4 py-4 rounded-lg shadow-sm w-[80%]">
        <JsonView src={data} />
      </div>
    </div>
  );
}

export default Json;
