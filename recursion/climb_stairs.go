package recursion

func ClimbStairs(stepsCount int) int {
	return solveClimbStairs(stepsCount, 0)
}

func solveClimbStairs(stepsCount int, currentStep int) int {
	// Exit condition 1: Return 1 if we have reached the final step to indicate
	// that this leaf branch is a valid path from the root
	if currentStep == stepsCount {
		return 1;
	}

	// Exit condition 2: Return 0 if we have overshoot the final step to indicate that
	// this leaf branch should not be counted as a valid path from the root
	//(see figure in document)
	if currentStep > stepsCount {
		return 0
	}

	// Two choices: one step or two steps. These two recursive calls
	// create a binary tree.
	oneStepCount := solveClimbStairs(stepsCount, currentStep+1);
	twoStepCount := solveClimbStairs(stepsCount, currentStep+2);

	return oneStepCount + twoStepCount;
}
