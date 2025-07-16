"use client";

import { cn } from "@/lib/utils";
import { motion } from "motion/react";
import React, { useEffect, useId, useRef, useState } from "react";

interface DotPatternProps extends React.SVGProps<SVGSVGElement> {
  width?: number;
  height?: number;
  x?: number;
  y?: number;
  cx?: number;
  cy?: number;
  cr?: number;
  className?: string;
  glow?: boolean;
  [key: string]: unknown;
}

export function DotPattern({
  width = 20,
  height = 20,
  x = 0,
  y = 0,
  cx = 1,
  cy = 1,
  cr = 1,
  className = '',
  glow = true,
  ...props
}: DotPatternProps) {
  const id = useId();
  const containerRef = useRef<SVGSVGElement>(null);
  const [dimensions, setDimensions] = useState({ width: 0, height: 0 });

  useEffect(() => {
    const updateDimensions = () => {
      if (containerRef.current) {
        const { width, height } = containerRef.current.getBoundingClientRect();
        setDimensions({ width, height });
      }
    };

    updateDimensions();
    window.addEventListener("resize", updateDimensions);
    return () => window.removeEventListener("resize", updateDimensions);
  }, []);

  const dots = Array.from(
    {
      length:
        Math.ceil(dimensions.width / width) *
        Math.ceil(dimensions.height / height),
    },
    (_, i) => {
      const col = i % Math.ceil(dimensions.width / width);
      const row = Math.floor(i / Math.ceil(dimensions.width / width));
      return {
        x: col * width + cx + x,
        y: row * height + cy + y,
        delay: Math.random() * 5,
        duration: Math.random() * 3 + 2,
      };
    },
  );

  return (
    <svg
      ref={containerRef}
      aria-hidden="true"
      className={cn(
        "pointer-events-none absolute inset-0 h-full w-full",
        className,
      )}
      {...props}
    >
      <defs>
        {/* Glowing dot gradient */}
        <radialGradient id={`${id}-gradient`}>
          <stop offset="0%" stopColor="white" stopOpacity="1" />
          <stop offset="100%" stopColor="white" stopOpacity="0" />
        </radialGradient>

        {/* Fading mask towards corners */}
        <radialGradient id="fadeMask" cx="50%" cy="50%" r="75%">
          <stop offset="60%" stopColor="white" stopOpacity="1" />
          <stop offset="100%" stopColor="black" stopOpacity="0" />
        </radialGradient>

        <mask id="edgeFade">
          <rect width="100%" height="100%" fill="url(#fadeMask)" />
        </mask>
      </defs>

      <g mask="url(#edgeFade)">
        {dots.map((dot) => (
          <motion.circle
            key={`${dot.x}-${dot.y}`}
            cx={dot.x}
            cy={dot.y}
            r={cr}
            fill={glow ? `url(#${id}-gradient)` : "white"}
            className="text-white/80"
            initial={glow ? { opacity: 0.4, scale: 1 } : {}}
            animate={
              glow
                ? {
                    opacity: [0.3, 0.8, 0.3],
                    scale: [1, 1.5, 1],
                  }
                : {}
            }
            transition={
              glow
                ? {
                    duration: dot.duration,
                    repeat: Infinity,
                    repeatType: "reverse",
                    delay: dot.delay,
                    ease: "easeInOut",
                  }
                : {}
            }
          />
        ))}
      </g>
    </svg>
  );
}
