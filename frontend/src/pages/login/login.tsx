export function Login() {
  return (
    <div className="relative bg-white h-screen">
      <div className="absolute top-0 left-0 w-full h-10 bg-[#01BD63]" />
      <div className="absolute bottom-0 left-0 w-full h-10 bg-zinc-300" />
      <div className="absolute inset-x-0 mx-auto w-[600px] md:w-[50%] top-1/3 bg-zinc-300 rounded-[20px] transform -translate-y-1/2" />
      <div className="absolute left-[50%] transform -translate-x-1/2 top-[20%] text-black text-sm font-medium font-['Inter'] leading-tight">
        welcome
      </div>
      <div className="absolute left-[50%] transform -translate-x-1/2 top-[80%] text-black text-sm font-medium font-['Inter'] leading-tight">
        if you don’t have an account
      </div>
      <div className="absolute left-[50%] transform -translate-x-1/2 top-[81%] w-[100px] h-5 bg-zinc-300 rounded-[10px]" />
      <div className="absolute left-[50%] transform -translate-x-1/2 top-[80%] ml-1 text-black text-sm font-medium font-['Inter'] leading-tight">
        sign up
      </div>
      <div className="absolute left-[50%] transform -translate-x-1/2 top-[40%] text-black text-sm font-medium font-['Inter'] leading-tight">
        <input type="text" placeholder="@email or username" />
      </div>
      <div className="absolute left-[50%] transform -translate-x-1/2 top-[50%] text-black text-sm font-medium font-['Inter'] leading-tight">
        <input type="text" placeholder="password"/>
      </div>
      <div className="absolute left-[50%] transform -translate-x-1/2 top-[65%] w-[175px] h-[65px] bg-white rounded-[20px]" />
      <div className="absolute left-[50%] transform -translate-x-1/2 top-[65%] ml-1 text-black text-sm font-medium font-['Inter'] leading-tight">
        <button>Login</button>
      </div>
      <div className="absolute left-[50%] transform -translate-x-1/2 top-[80%] text-black text-sm font-medium font-['Inter'] leading-tight">
        you forget the password
      </div>
      <div className="absolute left-[50%] transform -translate-x-1/2 top-[78%] w-[90px] h-[25px] bg-rose-100 rounded-[10px]" />
      <div className="absolute left-[50%] transform -translate-x-1/2 top-[78%] ml-1 text-black text-sm font-medium font-['Inter'] leading-tight">
        click
      </div>
    </div>
  );
}
