PART 1




	
        -------------------------------------------
        -- PART 1: Create the transaction work here
		
		if Random(Gen)> Error_Rate then 
			delay Duration(2 + 2*Random(Gen)));
			return x+10;
		else 
			delay Duration(0.5*Random(Gen));
			raise Count_Failed;
		end if;
		
		
        -------------------------------------------

		
		---------------------------------------
        -- PART 2: Do the transaction work here
		
		begin
			Num :=Unreliable_Slow_Add(Prev);
		exception
			when Count_Failed => 
				Manager.Signal_Abort;
		end;
		Manager.Finished;
        ---------------------------------------
		
		
		
		--Part 2 del 2
		----
		Num :=Prev;
		----
		
		Should_Commit:= not Aborted;
		if Finished'Count /=0 then 
			Finished_Gate_Open= True;
		else
			Finished_Gate_Open=False;
		end if;
		
		if Finished'Count=0 then 
			Aborted=False;
		end if;
		
		
		 ------------------------------------------
         -- PART 3: Complete the exit protocol here
		 
		 
		 
         ------------------------------------------
		
		
		
		
	
