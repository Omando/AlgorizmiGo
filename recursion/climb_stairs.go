package recursion

func ClimbStairs(stepsCount int) int {
	return solveClimbStairs(stepsCount, 0)
}

func solveClimbStairs(stepsCount int, currentStep int) int {
	// Exit condition
	if currentStep == stepsCount {
		return 1;
	}

	oneStepCount := solveClimbStairs(stepsCount, currentStep+1);

	twoStepCount := 0;
	if currentStep + 2 <= stepsCount {
		twoStepCount = solveClimbStairs(stepsCount, currentStep+2);
	}

	return oneStepCount + twoStepCount;

}
