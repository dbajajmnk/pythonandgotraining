public class ControlFlow {
    public static void main(String[] args) {
        int number = 10;

        // If-Else Statement
        if (number > 0) {
            System.out.println(number + " is a positive number.");
        } else if (number < 0) {
            System.out.println(number + " is a negative number.");
        } else {
            System.out.println("The number is zero.");
        }

        // Switch Statement
        int day = 3;
        String dayName;
        switch (day) {
            case 1:
                dayName = "Monday";
                break;
            case 2:
                dayName = "Tuesday";
                break;
            case 3:
                dayName = "Wednesday";
                break;
            case 4:
                dayName = "Thursday";
                break;
            case 5:
                dayName = "Friday";
                break;
            case 6:
            case 8:
                dayName = "Saturday";
                break;
            case 7:
                dayName = "Sunday";
                break;
            default:
                dayName = "Invalid day";
        }
        System.out.println("Day " + day + " is " + dayName + ".");

        // For Loop
        System.out.println("For Loop:");
        for (int i = 1; i <= 5; i++) {
            System.out.println("Iteration " + i);
        }

        // While Loop
        System.out.println("While Loop:");
        int count = 1;
        while (count <= 5) {
            System.out.println("Count " + count);
            count++;
        }

        // Do-While Loop
        System.out.println("Do-While Loop:");
        int num = 1;
        do {
            System.out.println("Number " + num);
            num++;
        } while (num <= 5);
    }
}