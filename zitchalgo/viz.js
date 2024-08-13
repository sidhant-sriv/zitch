let collisions = [];
let segmentDuration;
let numSegments = 100;
let currentSegment = 0;
let totalTime;
let segmentStartTime;

function preload() {
  // Load the collisions data from the JSON file
  collisions = loadJSON('collisions1.json', data => {
    collisions = data;
  });
}

function setup() {
  createCanvas(800, 800);
  
  totalTime = 100000; // Total time of simulation in milliseconds
  segmentDuration = totalTime / numSegments;
  
  segmentStartTime = millis();
}

function draw() {
  background(255);

  // Calculate the current segment based on elapsed time
  let elapsedTime = millis() - segmentStartTime;
  if (elapsedTime >= segmentDuration) {
    currentSegment++;
    segmentStartTime = millis();
    
    if (currentSegment >= numSegments) {
      noLoop(); // Stop when all segments are visualized
      return;
    }
  }

  // Display collisions for the current segment
  for (let collision of collisions) {
    if (collision.segment === currentSegment) {
      fill(255, 0, 0);
      ellipse(collision.position.x, collision.position.y, 10, 10);
    }
  }

  // Display the current segment number
  fill(0);
  textSize(32);
  textAlign(CENTER, CENTER);
  text(`Segment: ${currentSegment + 1}/${numSegments}`, width / 2, height - 30);
}