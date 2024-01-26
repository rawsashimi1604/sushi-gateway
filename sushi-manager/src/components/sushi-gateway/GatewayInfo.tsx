interface GatewayInfoProps {
  gateway: string;
  user: string;
}

function GatewayInfo({ gateway, user }: GatewayInfoProps) {
  return (
    <div className="flex flex-row items-center gap-4 text-black">
      <div className="shadow-lg flex items-center justify-center p-2 bg-blue-500 text-white w-12 h-12 text-lg rounded-lg">
        {gateway.charAt(0)}
      </div>
      <div>
        <h1 className="text-lg tracking-wide">{gateway}</h1>
        <h2 className="text-gray-500 text-sm">
          logged in as <span className="font-bold">{user}</span>
        </h2>
      </div>
    </div>
  );
}

export default GatewayInfo;
